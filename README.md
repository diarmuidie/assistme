# Sony GPS Ephemeris Data Download Tool

This is a simple tool to download the latest GPS Ephemeris data (GPS prediction or assistme file) onto a Sony Cyber-shot camera (tested on the Sony DSC HX90V) on a Mac computer.

This tool is needed because the [Sony PlayMemories software](http://support.d-imaging.sony.co.jp/www/disoft/int/download/playmemories-home/mac/en/), which is normally responsible for updating the GPS data, is [not yet compatible with macOS 10.13](http://sony-eur-eu-en-web--eur.custhelp.com/app/answers/detail/a_id/143062/~/macos-10.13-%28high-sierra%29-compatibility-information-for-application-software).

## Usage
Download a compiled binary or install and build:
```
go get github.com/diarmuidie/assistme
go install
```

Run the script passing in the mount location of the SD card or camera (in this example the SD card is called `Untitled` and is mounted in the `/Volumes` folder, where MacOS mounts SD cards by default):
```
assistme /Volumes/Untitled
```
