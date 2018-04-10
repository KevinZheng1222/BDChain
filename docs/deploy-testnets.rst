Deploy a Testnet
================

Now that we've seen how ABCI works, and even played with a few
applications on a single validator node, it's time to deploy a test
network to four validator nodes. For this deployment, we'll use the
``basecoin`` application.

Manual Deployments
------------------

It's relatively easy to setup a bdc cluster manually. The only
requirements for a particular bdc node are a private key for the
validator, stored as ``priv_validator.json``, and a list of the public
keys of all validators, stored as ``genesis.json``. These files should
be stored in ``~/.bdc/config``, or wherever the ``$TMHOME`` variable
might be set to.

Here are the steps to setting up a testnet manually:

1) Provision nodes on your cloud provider of choice
2) Install bdc and the application of interest on all nodes
3) Generate a private key for each validator using
   ``bdc gen_validator``
4) Compile a list of public keys for each validator into a
   ``genesis.json`` file.
5) Run ``bdc node --p2p.persistent_peers=< peer addresses >`` on each node,
   where ``< peer addresses >`` is a comma separated list of the IP:PORT
   combination for each node. The default port for bdc is
   ``46656``. Thus, if the IP addresses of your nodes were
   ``192.168.0.1, 192.168.0.2, 192.168.0.3, 192.168.0.4``, the command
   would look like:
   ``bdc node --p2p.persistent_peers=192.168.0.1:46656,192.168.0.2:46656,192.168.0.3:46656,192.168.0.4:46656``.

After a few seconds, all the nodes should connect to each other and start
making blocks! For more information, see the bdc Networks section
of `the guide to using bdc <using-bdc.html>`__.

Automated Deployments
---------------------

While the manual deployment is easy enough, an automated deployment is
usually quicker. The below examples show different tools that can be used
for automated deployments.

Automated Deployment using Kubernetes
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The `mintnet-kubernetes tool <https://github.com/bdc/tools/tree/master/mintnet-kubernetes>`__
allows automating the deployment of a bdc network on an already
provisioned Kubernetes cluster. For simple provisioning of a Kubernetes
cluster, check out the `Google Cloud Platform <https://cloud.google.com/>`__.

Automated Deployment using Terraform and Ansible
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The `terraform-digitalocean tool <https://github.com/bdc/tools/tree/master/terraform-digitalocean>`__
allows creating a set of servers on the DigitalOcean cloud.

The `ansible playbooks <https://github.com/bdc/tools/tree/master/ansible>`__
allow creating and managing a ``basecoin`` or ``ethermint`` testnet on provisioned servers.

Package Deployment on Linux for developers
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The ``bdc`` and ``basecoin`` applications can be installed from RPM or DEB packages on
Linux machines for development purposes. The packages are configured to be validators on the
one-node network that the machine represents. The services are not started after installation,
this way giving an opportunity to reconfigure the applications before starting.

The Ansible playbooks in the previous section use this repository to install ``basecoin``.
After installation, additional steps are executed to make sure that the multi-node testnet has
the right configuration before start.

Install from the CentOS/RedHat repository:

::

    rpm --import https://bdc-packages.interblock.io/centos/7/os/x86_64/RPM-GPG-KEY-bdc
    wget -O /etc/yum.repos.d/bdc.repo https://bdc-packages.interblock.io/centos/7/os/x86_64/bdc.repo
    yum install basecoin

Install from the Debian/Ubuntu repository:

::

    wget -O - https://bdc-packages.interblock.io/centos/7/os/x86_64/RPM-GPG-KEY-bdc | apt-key add -
    wget -O /etc/apt/sources.list.d/bdc.list https://bdc-packages.interblock.io/debian/bdc.list
    apt-get update && apt-get install basecoin

