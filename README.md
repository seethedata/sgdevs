sgdevs
---

sgdevs takes a text file containing a list of VMAX storage groups and produces a file containing all of the devices in those storage groups in a format suitable for use as a filter file in SymmMerge.

The name of the file for the list of storage groups is expected to be `storagegroups.txt` and the config information is read from a `symapi_db.bin` file.

Solutions Enabler `v7.4.0` or newer is required

*Note: Currently the output file uses UNIX-style newlines, which may make SymmMerge on Windows cranky. You may convert the newlines with the tool of your choosing, e.g. ux2dos, Notepad++, etc.  *