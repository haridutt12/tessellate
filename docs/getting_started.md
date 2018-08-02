# Tessellate: Getting Started.

# Pre requisite:
Download the latest Tessellate CLI binary from: https://github.com/tsocial/tessellate/releases 
Download the binary as per your Operating System: Linux or Mac.

# Terminologies used

### 1. Workspace
A workspace identifies with the environment in which you're deploying your infrastructure.
This workspace could be a dev, staging or a production environment.

- __Creating a new Workspace:__
Ideally, one must reuse the workspaces already created by us: `dev`, `staging` or `production`.
But, if you're working on a provider apart from aws, it would be best to create a separate workspace for every provider.
Ex: alicloud can have workspaces named: `alicloud-dev`, `alicloud-staging` or `alicloud-production`.

We need to feed your provider details while creating the workspace. An ideal providers.tf.json would look as follows:

[providers.tf.json]:
```
{"variable":{"region":{}},"provider":[{"alicloud":{"access_key":"AKIAJ6APYYH6OXBRBDXQ","region":"${var.region}","secret_key":"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"}}]}
```
__What is important to note here, is that a value for var.region is expected before one runs any layout.__
__Do not redefine the region variable, just pass a value for region.__

```
$ ./tessellate_cli workspace create -w alicloud-dev -p providers.tf.json -a control.ha.tsengineering.io:9977
```

- __Get the workspace details:__
To check if you're workspace was created successfully, execute:
```
$ ./tessellate_cli workspace get -w alicloud-dev -a control.ha.tsengineering.io:9977
```

### 2. Layout.

A layout basically identifies with the terraform scripts that you have written for deploying the infrastructure.

Let's say you have a directory folder named: __"/home/username/my_tf_scripts"__

in my_tf_scripts/ you have all your tf.json files: which are responsible for bringing up different instances and resources.

#### STEP 1: Creating a Layout:

```
$ ./tessellate_cli layout create -w staging -l my-layout -d /home/username/my_tf_scripts -a control.ha.tsengineering.io:9977
```

#### STEP 2: Get the Layout:

To check if your layout was created successfully, execute:

```
$ ./tessellate_cli layout get -w staging -l my-layout -a control.ha.tsengineering.io:9977
```

#### STEP 3: Apply the layout:

Here we will run a `terraform apply`.

But let's say before you do an apply, you want to do a dry run, and see `terraform plan`'s output.
Execute:
```
$ ./tessellate_cli layout apply -w staging -l my-layout -v /home/username/variables.tfvars.json -a control.ha.tsengineering.io:9977 --dry
```

The `dry` flag will run a `terraform plan` for you.
When you're ready to apply, remove the `dry` flag in the above command.


#### STEP 4: Destroy the layout:

To destroy the layout that you created, execute:

```
$ ./tessellate_cli destroy -w staging -l my-layout -v /home/username/variables.tfvars.json -a control.ha.tsengineering.io:9977
```

