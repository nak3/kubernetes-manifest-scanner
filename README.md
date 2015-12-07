kubernetes-manifest-scanner
==============
kubernetes-manifest-scanner is CLI to refer manifest file from swagger API

![kubernetes-manifest-scanner](https://raw.githubusercontent.com/wiki/nak3/kubernetes-manifest-scanner/gif/kms-command.gif)

Quick start
-----

##### 1. Build

~~~
git clone https://github.com/nak3/kubernetes-manifest-scanner.git
cd kubernetes-manifest-scanner
./build.sh
~~~

##### 2. Setup path and bash completion

~~~
export PATH="${PWD}/bin/:$PATH"
source bash-comp/kuberenetes-manifest-scanner
~~~

##### 3. Now, you can use it!

e.g.
~~~
kubernetes-manifest-scanner sample v1.Pod -d 20
~~~

NOTE: You `-f` flag will point to kubernetes's [upstream swagger](https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/swagger-spec/v1.json) by default.

Command usage
----

##### kubernetes-manifest-scanner sample

Output manifest with all parameteres

~~~
kubernetes-manifest-scanner -k -f https://ose3-master.example.com:8443/swaggerapi/api/v1 -d 20 sample v1.Node
{
	"apiVersion": "string",
	"kind": "string",
	"metadata": {
		"annotations": {
			"any": "e.g. label"
		},
		"creationTimestamp": "string",
		"deletionGracePeriodSeconds": "integer",
		"deletionTimestamp": "string",
		"generateName": "string",
		"generation": "integer",
		"labels": {
			"any": "e.g. label"
		},
		"name": "string",
		"namespace": "string",
		"resourceVersion": "string",
		"selfLink": "string",
		"uid": "string"
	},
	"spec": {
		"externalID": "string",
		"podCIDR": "string",
		"providerID": "string",
		"unschedulable": "boolean"
	}
}

~~~


| Name                       | Description                                             |
|:---------------------------|:--------------------------------------------------------|
|`-f` *swagger URL/file*     | Path to the swagger URL or JSON file .                  |
|`-k`                        | Set with `-f https://...` and allow insecure access.    |
|`-d` *depth*                | Some parameters set as `$ref` and expand them with the depth.|


##### kubernetes-manifest-scanner snippet

Refer a configuration parameter as snippet

~~~
kubernetes-manifest-scanner snippet deletionGracePeriodSeconds
{
	"deletionGracePeriodSeconds": {
		"description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
		"format": "int64",
		"type": "integer"
	}
}
~~~

| Name                       | Description                                             |
|:---------------------------|:--------------------------------------------------------|
|`-f` *swagger URL/file*     | Path to the swagger URL or JSON file .                  |
|`-k`                        | Set with `-f https://...` and allow insecure access.    |


##### kubernetes-manifest-scanner itemlist

Get items to path `kubernetes-manifest-scanner sample <ITEM>`, but you don't need to use this subcommand if you enabled bash-completion.
