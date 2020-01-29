ipxe\_server
============

Description
-----------

Assuming you have your own DHCP/TFTP/HTTP server(s)...

This is a simple script to serve .ipxe scripts once and then delete them.
Useful if you want to chainload ipxe scripts to "break the chain" for
small dev-test like setups.

Installation
------------

1. Setup your DHCP and TFP server for network booting. Goal is that each
 time your machines will network boot to an ipxe loader. The current
 setup assumes you are doing static DHCP address assignment (it relies
 on IP addresses to associate machines with ipxe scripts).

2. Setup your [ipxe chainloading](http://ipxe.org/howto/chainloading) boot
program.

 I have some older firware machines so I built a custom undionly.kpxe boot
 program. I choose to EMBED my ipxe script as it made my setup cleaner
 and I didn't want to edit my DHCP setup scripts other than provide
 a custom boot file. Like this:

```bash
cd ipxe/src
make bin/undionly.kpxe NO_WERROR=1 V=1 EMBED=ocp.ipxe
```

Where ocp.ipxe contains something like this:

```
#!ipxe

echo Here we go...
dhcp net0 && route
chain --replace --autofree http://192.168.130.10:10080?mac=${net0/mac}
```

3. Checkout this project. Install golang.

4. Create softlinks to the ipxe files you want to run custom ipxe scripts.
Example:

```bash
ln -s scripts/coreos/bootstrap.ipxe 172.19.0.6
```

5. Start the server on port (defaults to 10080):
```bash
go run src/ipxe.go
```

How it works
------------

ipxe.go listens for HTTP requests on 10080. If then serves the contents of
any file named after the IP address of the incoming HTTP request and removes
that file (thus breaking the reboot forever loop). If no file is
found a simple 'exit' ipxe script is served which should allow the machine
to boot normally.

Anytime you want to re-image a machine create a simple soft link from the
IP of the machine you want to re-image to the ipxe file you want to boot.
