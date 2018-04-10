Install bdc
==================

From Binary
-----------

To download pre-built binaries, see the `Download page <https://bdc.com/download>`__.

From Source
-----------

You'll need ``go``, maybe `dep <https://github.com/golang/dep>`__, and the bdc source code.

Install Go
^^^^^^^^^^

Make sure you have `installed Go <https://golang.org/doc/install>`__ and
set the ``GOPATH``. You should also put ``GOPATH/bin`` on your ``PATH``.

Get Source Code
^^^^^^^^^^^^^^^

You should be able to install the latest with a simple

::

    go get github.com/bdc/bdc/cmd/bdc

Run ``bdc --help`` and ``bdc version`` to ensure your
installation worked.

If the installation failed, a dependency may have been updated and become
incompatible with the latest bdc master branch. We solve this
using the ``dep`` tool for dependency management.

First, install ``dep``:

::

    make get_tools

Now we can fetch the correct versions of each dependency by running:

::

    cd $GOPATH/src/github.com/bdc/bdc
    make get_vendor_deps
    make install

Note that even though ``go get`` originally failed, the repository was
still cloned to the correct location in the ``$GOPATH``.

The latest bdc Core version is now installed.

Reinstall
---------

If you already have bdc installed, and you make updates, simply

::

    cd $GOPATH/src/github.com/bdc/bdc
    make install

To upgrade, there are a few options:

-  set a new ``$GOPATH`` and run
   ``go get github.com/bdc/bdc/cmd/bdc``. This
   makes a fresh copy of everything for the new version.
-  run ``go get -u github.com/bdc/bdc/cmd/bdc``,
   where the ``-u`` fetches the latest updates for the repository and
   its dependencies
-  fetch and checkout the latest master branch in
   ``$GOPATH/src/github.com/bdc/bdc``, and then run
   ``make get_vendor_deps && make install`` as above.

Note the first two options should usually work, but may fail. If they
do, use ``dep``, as above:

::

    cd $GOPATH/src/github.com/bdc/bdc
    make get_vendor_deps
    make install

Since the third option just uses ``dep`` right away, it should always
work.

Troubleshooting
---------------

If ``go get`` failing bothers you, fetch the code using ``git``:

::

    mkdir -p $GOPATH/src/github.com/bdc
    git clone https://github.com/bdc/bdc $GOPATH/src/github.com/bdc/bdc
    cd $GOPATH/src/github.com/bdc/bdc
    make get_vendor_deps
    make install

Run
^^^

To start a one-node blockchain with a simple in-process application:

::

    bdc init
    bdc node --proxy_app=kvstore
