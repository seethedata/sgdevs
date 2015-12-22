sgdevs
---

sgdevs takes a text file containing a list of VMAX storage groups and produces a file containing all of the devices in those storage groups in a format suitable for use as a filter file in SymmMerge.

Requirements:
* A file with the name `storagegroups.txt` in the same directory as this program . The format of the file is one storage group name per line. 
* A `symapi_db.bin` file in the same directory as this program that can be used to read VMAX config information.
* Solutions Enabler `v7.4.0` or newer

*Note: Currently the output file uses UNIX-style newlines, which may make SymmMerge on Windows cranky. You may convert the newlines with the tool of your choosing, e.g. ux2dos, Notepad++, etc.*