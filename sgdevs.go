//Package symmsummary pulls out disk info from symmapi_db.bin.

package main

import (
	"bufio"
	"fmt"
	"github.com/seethedata/symmtools"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"strconv"
)

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

func main() {
	fileName := "symapi_db.bin"
	os.Setenv("SYMCLI_OFFLINE", "1")
	os.Setenv("SYMCLI_DB_FILE", fileName)

	devPattern := regexp.MustCompile("^[0-9a-fA-F]{4,5}\\b")
	requiredVersion := "V7.4.0"
	cfgexe := symmtools.LocateFile("symcfg.exe")
	devexe := symmtools.LocateFile("symdev.exe")
	version := symmtools.GetVersion(cfgexe)
	if requiredVersion > version {
		log.Fatal(requiredVersion + " is required. Installed version is " + version)
	}

	storageGroups := make([]string, 0)

	devfile, err := os.Open("storagegroups.txt")
	check("Opening File: ", err)
	nextLine := bufio.NewScanner(devfile)
	for nextLine.Scan() {
		storageGroups = append(storageGroups, nextLine.Text())
	}

	label := regexp.MustCompile("(DMX|VMAX)")
	sids := make([]string, 0)

	command := exec.Command(cfgexe, "list")
	stdout, err := command.StdoutPipe()
	check("Getting Symm list: ", err)
	command.Start()
	result := bufio.NewScanner(stdout)
	check("Command Output: ", err)

	for result.Scan() {
		if label.MatchString(result.Text()) == true {
			arrayData := strings.Fields(result.Text())
			sids = append(sids, arrayData[0])
		}
	}

	for _, sid := range sids {
		fmt.Printf("Extracting selected devices for SID %s...", sid)
		outfile, err := os.Create(sid + "-filteredDevices.txt")
		sizefile, err:=os.Create(sid + "-storageGroupSizes.csv")
		_,err=sizefile.WriteString(fmt.Sprintf("%s,%s\n","Storage Group", "Size(MB)"))
		check("Creating file: ", err)
		sgList := make([]string, 0)
		for _, sg := range storageGroups {
			sgSize:=0
			command := exec.Command(devexe, "list", "-sid", sid, "-sg", sg)
			stdout, err := command.StdoutPipe()
			check("Command Output: ", err)
			command.Start()
			result := bufio.NewScanner(stdout)
			for result.Scan() {
				devLine := result.Text()
				deviceSize:=0
				if devPattern.MatchString(devLine) == true {
					arrayData := strings.Fields(devLine)
					device := arrayData[0]
					deviceSize,err=strconv.Atoi(arrayData[len(arrayData)-1])
					sgList = append(sgList, device)
					_, err := outfile.WriteString(fmt.Sprintf("%s,%s\n", device, sid))
					check("Writing to file: ", err)
					sgSize=sgSize +  deviceSize
					check("Convert string to int: ",err)
				}
				
			}
			_, err = sizefile.WriteString(fmt.Sprintf("%s,%d\n",sg,sgSize))
			check("Writing size file: ",err)
		}
		fmt.Printf("Done.\n")
		fmt.Printf("File created:%s with %d devices.\n", outfile.Name(), len(sgList))
		fmt.Printf("File created:%s.\n", sizefile.Name())
	}

}
