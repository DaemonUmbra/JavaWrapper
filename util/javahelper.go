/**
 * Minecraft Forge
 * Copyright (c) 2016.
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation version 2.1
 * of the License.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
 */
package util

import (
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

const installerVersion = "@VERSION@"


func IsJavaInstalled() bool {
	jv, err := exec.LookPath("java")

	if err != nil {
		color.Red("Java has not been found!")
		return false
	} else {
		color.Green("Java has been found at: %s", jv)
		return true
	}
}

//I'm going to need a CI for this one
func InstallForge()  {
	var url = "http://files.minecraftforge.net/maven/net/minecraftforge/forge/"+ installerVersion +
		"/forge-" + installerVersion + "-installer.jar"
	color.Green("Downloading forge")
	DownloadFromUrl(url, getMcDir())
	color.Yellow("Running forge " +installerVersion + "-installer.jar")
	_, err := exec.Command("java", "-jar", getMcDir() + "/forge-" +installerVersion + "-installer.jar").CombinedOutput()

	if err != nil {
		color.Red("There was a problem running the installer %s", err)
	}

	color.Yellow("removing installer")
	os.Remove(getMcDir() +"/forge-" + installerVersion + "-installer.jar")
	os.Remove(getMcDir() +"/forge-" + installerVersion + "-installer.jar.log")

}

func LaunchWithSysJava() {
	color.Yellow("ForgeWrapper is now lauching the laucnher with system JRE")

	out, err := exec.Command("java", "-jar", getMcDir()+"/launcher.jar").CombinedOutput()

	if err != nil {
		println(err)
	}

	println(out)
}

func LaunchWithMojangJava() {
	darwinJRE := getRuntimeJREDir() + "/bin/java"
	winJRE := getRuntimeJREDir() + "/bin/java.exe"
	color.Yellow("Now running the Launcher from Mojang JRE")

	if GetThisPlatform() == "windows" {
		exec.Command(winJRE, "-jar", getMcDir()+"/launcher.jar").Run()
	} else if GetThisPlatform() == "darwin" {
		exec.Command(darwinJRE, "-jar", getMcDir()+"/launcher.jar").Run()
	} else if GetThisPlatform() == "linux" {
		color.Red("Sorry Mojang has not built a JRE for linux please download from go to " +
			"http://openjdk.java.net/install/ or " +
			"http://www.oracle.com/technetwork/java/javase/downloads/index.html to download the latest java 8.")
		os.Exit(3)
	}

}

func IsJavaVersionValid() bool {
	out, _ := exec.Command("java", "-version").CombinedOutput()

	if strings.Contains(string(out), "1.8") {
		return true
	} else {
		return false
	}
}

func JreLauncher() {
	if IsValidPlatFrom() {
		checkForMcdir()
		CheckForLauncher()
		if IsJavaVersionValid() {
			LaunchWithSysJava()
		} else {
			checkForRuntime()
			LaunchWithMojangJava()
		}
	}
}

func ModedLauncher()  {
	if installerVersion != "@VERSION@" {
		checkForMcdir()
		InstallForge()
	} else {
		JreLauncher()
	}
}
