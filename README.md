# Cryptseaside

![sample.png](./README_image/2021_11_12T07_22_48_142355004_09_00.png)

Welcome to the Cryptseaside project.
Cryptseaside generates seaside images using Unix nanoseconds as the seed value.

## Usage

```shell
go install github.com/rhymof/cryptseaside/cmd/cryptseaside
```

### options

```
  -i    save as png. needs a directory named 'images' (default true)
  -s    stdout seed (default true)
  -u    stdout URI (default true)
  -v    stdout version (default true)
```

## Current version

Now, Cryptseaside version is alpha.
Destructive changes may occur.

## Probability of generation

|Type|Value|Probability|
|:--|:--|:--|
|Time|day|56/256|
|Time|evening|100/256|
|Time|night|100/256|

## Sample images

![day.png](./README_image/2021_11_12T07_20_44_687607499_09_00.png)
![evening.png](./README_image/2021_11_12T07_22_36_313230478_09_00.png)

## Todo(alpha)

* [ ] a short range view objects(new package)
* [ ] Feature: plot object(internal)
* [ ] clouds
* [ ] Special mode(new package)
* [ ] Make sunset more beautiful