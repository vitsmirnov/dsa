"""

Input: button_count, emoji_count
Output: all variations

Example:
Input: button_count = 2, emoji_count = 3
Output:
|o|oo|
|oo|o|

"""


def main() -> None:
    button_count, emoji_count = 2, 3
    button_count, emoji_count = 3, 5
    button_count, emoji_count = 5, 7

    print(f"button_count = {button_count}, emoji_count = {emoji_count}")
    for version in distribute_emoji(button_count, emoji_count):
        print(version)


def distribute_emoji(button_count: int, emoji_count: int) -> list[str]:
    res = list()

    def _distribute_emoji(button_number: int, cur_emoji_count: int, cur_emoji: str) -> None:
        if button_number == button_count-1:
            res.append(cur_emoji + "o"*(emoji_count-cur_emoji_count)  + "|")
            return

        cur_max_emoji = emoji_count - cur_emoji_count - (button_count - button_number - 1)
        for _emoji_count in range(1, cur_max_emoji+1):
            _distribute_emoji(
                button_number+1,
                cur_emoji_count+_emoji_count,
                cur_emoji + "o"*_emoji_count + "|")

    _distribute_emoji(0, 0, "|")

    return res


if __name__ == "__main__":
    main()
