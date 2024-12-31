# doubutsu-uranai

The most accurate animal fortune-telling, your personality checker based on your date of birth.
Don't believe, me? Try it out!

https://ja.wikipedia.org/wiki/動物占い

From 1800 to 2024

# Install

```sh
```

# Usage

```sh
doubutsu-uranai

Enter your gender (M/F): M
Enter your date of birth (YYYY-MM-DD): 1982-02-02
Results for Gender: M, Date of Birth: 1982-02-02

Your Aspiration Character: ライオン
    Description: Lion: Bold, courageous, and naturally commands respect.
Your Decision-Making Character: ゾウ
    Description: Elephant: Wise, dependable, and emotionally strong.
Your Core Character: 黒ひょう
    Description: Black Panther: Charismatic and stylish, thrives on individuality.
Your Surface Character: こじか
    Description: Fawn: Innocent and pure, a seeker of peace and affection.
Your Hidden Character: ひつじ
    Description: Sheep: Gentle, nurturing, and values harmony in groups.
```

# Development Notes

```sh
go run main.go
go test ./...
```

# Deployment Notes

Since the gold data is proprietary, it is not included in the repository.

```sh
git tag v1.0.0
./build.sh # manually upload to Github releases
```