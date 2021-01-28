# QLang Specification

<sup>Written by Tarith Jayasooriya</sup>

## Abstract

This paper contains the specification of the QLang XML language extension. This will cover every tag and their uses



## Tags

1. qPack
2. qSet
3. q
4. choice

### qPack

```xml
<qPack author="tarith" text="Maths test"></qPack>
```



qPack is array which contains ***1 or more*** qSet(s), A qPack is similar to an test paper. qPack should be the root of the xml file.

| Attribute       | Value  | Description                   |
| --------------- | ------ | ----------------------------- |
| author          | String | The author of the questions   |
| text (Required) | String | The heading of the test paper |

| Children |
| -------- |
| qSet     |

### qSet

qSet is a array that contains ***1 or more*** q(s), A qPack is a main question.    

| Attribute       | Value  | Description                      |
| --------------- | ------ | -------------------------------- |
| text (Required) | String | The heading of the main question |

### q

q is a array that contains ***1 or more*** choice(s)

| Attribute     | Value  | Description                        |
| ------------- | ------ | ---------------------------------- |
| text (author) | String | The heading of the branch question |

| Children |
| -------- |
| choice   |

### choice

choice is a empty element.

| Attribute | Value         | Description                          |
| --------- | ------------- | ------------------------------------ |
| correct   | (true\|false) | whether the choice is correct or not |
| value     | String        | The heading of the choice question   |

