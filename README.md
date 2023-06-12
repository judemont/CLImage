# CLI Image Displayer
Watch pictures without leaving your terminal !
 
A go script to display images in your terminal, with colored characters.

![image](https://github.com/judemont/CLI-Image-Displayer/assets/96385330/4f472588-3e60-40ab-bc58-9cafe849bb60)


## How to install it ?
1. Download the latest release at [github.com/judemont/Image-Displayer/releases/latest](https://github.com/judemont/Image-Displayer/releases/latest)
2. Extracts the file and open the folder in a terminal
3. Make is executable :
```bash
chmod +x  CliImageDisplayer
```
3. If you want to be able to use it everywhere as a simple command, you can add it to the `/bin/` directory :
```bash
sudo mv CliImageDisplayer /bin/
```
## How to use it (usage) :
```
usage: main [<flags>] <image path>


Flags:
      --[no-]help     Show context-sensitive help (also try --help-long and --help-man).
  -w, --width=50      Image width (in characters)
      --[no-]version  Show application version.

Args:
  <image path>  The path of the image.
```
### Example
```bash
CliImageDisplayer --width 100 robot.png
```
```bash
CliImageDisplayer  -w 100 images/robot.png
```
