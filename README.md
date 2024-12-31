# `doubutsu-uranai`

The most accurate animal fortune-telling, **your personality checke**r based statistics, just by your gender and date of birth.

**Don't believe, me? Try it out!**

https://ja.wikipedia.org/wiki/動物占い

From 1800 to 2024

# Install

```sh
curl -sL https://raw.githubusercontent.com/kevincobain2000/doubutsu-uranai/master/install.sh | sh
```

# Usage

```sh
./doubutsu-uranai

Enter your gender (M/F): F
Enter your date of birth (YYYY-MM-DD): 1990-09-09
Results for Gender: F, Date of Birth: 1990-09-09

Your 希望キャラ(Aspiration Character): ペガサス(Pegasus)
Description Pegasus: Dreamy, creative, and unbound by limitations. Pegasus individuals are imaginative and visionary, often thinking outside the box. They are idealistic and have a strong desire to explore new possibilities, unencumbered by conventional limits. Their creativity knows no bounds, and they are drawn to artistic or unconventional pursuits. However, their dreamy nature can sometimes make them seem detached or impractical, as they are more focused on their visions than the everyday realities of life.

Your 意思決定キャラ(Decision-Making Character): ライオン(Lion)
Description Lion: Bold, courageous, and naturally commands respect. Lions are natural-born leaders, full of energy, confidence, and determination. They are charismatic and have the ability to inspire and motivate others with their vision. People with the Lion personality tend to be ambitious and always strive for success. While they are generous and protective of their loved ones, they can sometimes be perceived as domineering or stubborn due to their strong will and desire for control.

Your 本質キャラ(Core Character): ひつじ(Sheep)
Description Sheep: Gentle, nurturing, and values harmony in groups. Sheep people are compassionate, kind-hearted, and always look for ways to help those around them. They are sensitive to the needs of others and excel in creating harmonious environments. While they are caring and selfless, they may sometimes struggle with asserting themselves, as they dislike conflict and confrontation. Their empathetic nature makes them excellent listeners and trusted friends.

Your 表面キャラ(Surface Character): 猿(Monkey)
Description Monkey: Playful, energetic, and highly intelligent. People born under the Monkey sign are curious, quick-witted, and often the life of the party. They thrive on intellectual challenges and are always looking for new ways to solve problems. Monkeys are sociable and love to interact with others, using their humor and intelligence to make a lasting impression. While they are highly adaptable, they can sometimes be seen as restless or unpredictable, as they seek constant stimulation and excitement.

Your 隠れキャラ(Hidden Character): ゾウ(Elephant)
Description Elephant: Wise, dependable, and emotionally strong. Elephants are deeply loyal and steadfast, always there for those they care about. They possess a calm and grounded presence, providing guidance to those in need. Their wisdom comes from experience, and they are known to approach challenges with patience and thoughtfulness. However, they may be seen as slow to act, as they prefer to carefully consider all options before making decisions. They are nurturing and protective, making them great caregivers.
```

# Development Notes

```sh
go run main.go
go test ./...
```

# Deployment Notes

Since the gold data is proprietary, it is not included in the repository.

```sh
git tag v1.0.x
git push origin v1.0.x
./build.sh

Github > releases > create new release > choose tag > upload binary
```