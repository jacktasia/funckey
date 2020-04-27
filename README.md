# funckey

`funckey` is a simple go cli program that wraps lower level linux cli apps that are commonly mapped to function keys on laptops.

I made this so I could easily map by `volume` and `brightness` function keys on my laptop to work when running [`monand`](https://xmonad.org/) and [`xmobar`](https://hackage.haskell.org/package/xmobar) (see config examples below). If you're using gnome/kde you probably don't need something like this.

## Setup / Install


`go install github.com/jacktasia/funckey`


### required packages

Ensure you have `pactl` installed

and

``/sys/class/backlight/intel_backlight/brightness` exists.

### brightness permissions

Your user needs to be able to edit the `brightness` file:

`chown jack:jack /sys/class/backlight/intel_backlight/brightness`

or

`sudo nano /etc/udev/rules.d/90-myrules.rules` and add a line with:

```
KERNEL=="intel_backlight", SUBSYSTEM=="backlight", RUN+="/usr/bin/find /sys/class/backlight/intel_backlight/ -type f -name brightness -exec chown jack:jack {} ; -exec chmod 666 {} ;"
```

## Commands

### Controlling Volume
* `funckey volume down`        Decrease system volume by ~10%
* `funckey volume up`          Increase system volume by ~10%
* `funckey volume toggle-mute` Toggle the mute volume
* `funckey volume mute`        Mute system volume
* `funckey volume unmute`      Un-mute system volume
* `funckey volume get-percent` Get the system volume
* `funckey volume set-percent N` Set the system volume by percent

### Controlling Brightness
* `funckey brightness down`        Decrease screen brightness by ~10%
* `funckey brightness up`          Increase screen brightness by ~10%
* `funckey brightness get-percent` Get the screen brightness
* `funckey brightness set-percent N` Set the screen brightness by percent
* `funckey brightness status` Set the screen brightness status


### xomand keybindings

A snippet of my `xmonad.hs` config using `import XMonad.Util.EZConfig`:

```haskell
  ("<XF86MonBrightnessDown>", spawn "funckey brightness down"),
  ("<XF86MonBrightnessUp>", spawn "funckey brightness up"),
  ("<XF86AudioLowerVolume>", spawn "funckey volume down"),
  ("<XF86AudioRaiseVolume>", spawn "funckey volume up"),
  ("<XF86AudioMute>", spawn "funckey volume toggle-mute")
```

### xobar status config

In the `commands` section of the config:

```haskell
        , Run Com "funckey" ["brightness", "get-percent"] "mybrightness" 2
        , Run Com "funckey" ["volume", "get-percent"] "myvolume" 2
```

then in the `template` you can have something like:

```
| Screen: %mybrightness% | Vol: %myvolume%
```

