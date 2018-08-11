package dispatcher

import (
	"fmt"
	"log"

	"github.com/flosch/pongo2"
	"github.com/hashicorp/nomad/api"
	"github.com/tsocial/tessellate/tmpl"
)

func (c *client) GetOrSetBackup(s string) error {
	nConfig := api.DefaultConfig()
	nConfig.Address = c.cfg.Address

	if c.cfg.Username != "" {
		nConfig.HttpAuth = &api.HttpBasicAuth{
			Username: c.cfg.Username,
			Password: c.cfg.Password,
		}
	}

	cl, err := api.NewClient(nConfig)
	if err != nil {
		log.Printf("error while creating nomad client: %+v", err)
		return err
	}

	job, _, err := cl.Jobs().Info(s, nil)
	// Replace the Job if not found OR existing Job was dead.
	if job == nil || *job.Status == "dead" {
		if err := c.startBackupJob(s); err != nil {
			return err
		}
	}

	return nil
}

func backupCmd(consulHost string) string {
	command := "ts=`date +%s`;ts+=_backup.snap;"
	subCommand := "consul snapshot save -http-addr="+ consulHost +" $ts"

	log.Println(command)
	log.Println(subCommand)

	return command + subCommand
}

func (c *client) startBackupJob(jobID string) error {
	// Create a nomad job using go template
	var tmplStr = fmt.Sprintf(`
job "{{ job_id }}" {
  datacenters = ["{{ datacenter }}"]
  type        = "batch"

  periodic {
    cron             = "*/2 * * * * *"
    prohibit_overlap = true
  }

  group "{{ job_id }}" {
    count = 1

    task "backup_job" {
      driver = "raw_exec"

		config {
			command = "bash"
			args = ["-exc", "%v"]
      }

      resources {
        cpu    = {{ cpu }}
        memory = {{ memory }}
      }
    }
  }
}
`, backupCmd(c.cfg.ConsulAddr))

	cfg := pongo2.Context{
		"job_id":     jobID,
		"datacenter": c.cfg.Datacenter,
		"cpu":        c.cfg.CPU,
		"memory":     c.cfg.Memory,
	}

	nomadJob, err := tmpl.Parse(tmplStr, cfg)
	if err != nil {
		log.Printf("error while job parsing: %+v", err)
		return err
	}

	log.Println(nomadJob)

	nConfig := api.DefaultConfig()
	nConfig.Address = c.cfg.Address

	if c.cfg.Username != "" {
		nConfig.HttpAuth = &api.HttpBasicAuth{
			Username: c.cfg.Username,
			Password: c.cfg.Password,
		}
	}

	cl, err := api.NewClient(nConfig)
	if err != nil {
		log.Printf("error while creating nomad client: %+v", err)
		return err
	}

	jobs := cl.Jobs()
	job, err := jobs.ParseHCL(nomadJob, true)
	if err != nil {
		log.Printf("error while parsing job hcl: %+v", err)
		return err
	}

	resp, _, err := jobs.Register(job, nil)
	if err != nil {
		log.Printf("error while registering nomad job: %+v", err)
		return err
	}

	log.Printf("successfully started the job: %+v", resp)
	return nil
}
