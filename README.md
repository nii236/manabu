# Manabu

I and many others have always had trouble finding good material and learning methods for Japanese.

This project is an attempt to combine the following:

- Anki's excellent spaced repetition algorithm
- Twitter's enormous amount of data
- Kagome/Mecab morphological analysis
- Allowing students to learn at their own pace, on material that **they are actually interested in**

This project uses `glide` for vendoring dependencies.

Docker support will be added soon to handle containerising Kagome.

A React WebUI and REST API will be added eventually to allow user interaction.

# Spinup

```
brew install glide
glide install
glide rebuild
go build
./gank
```
