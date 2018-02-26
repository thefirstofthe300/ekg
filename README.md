## Machine EKG
[![](https://build.danielseymour.info/api/badges/thefirstofthe300/ekg/status.svg)](https://build.danielseymour.info/thefirstofthe300/go-playground)

The machine EKG is a tool to be used to diagnose various issues with machine configuration. This tool was written with Google Compute Engine machines in mind,
although pull requests for other types of VMs (including bare-metal) are welcome.

Due to the fact that this tool digs deeply into the Linux's internals for much
of its information, I do not intend to support non-Linux VMs.

#### TODO

* Reading the machine's DHCP configuration (if possible???)
* Add support to run as a privileged Docker container
* Provide multiple methods to export and share the resulting file, e.g. Google
Cloud Storage, text file, Stackdriver Logging (???) 
