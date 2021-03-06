# ansi color pkg and tools

![image](https://user-images.githubusercontent.com/1940588/104885143-7eb07680-59a2-11eb-8749-cb8c3558b7ff.png)

## ANSI Escape Sequences

[ANSI Escape Sequences](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)

Standard escape codes are prefixed with `Escape`:

- Ctrl-Key: `^[`
- Octal: `\033`
- Unicode: `\u001b`
- Hexadecimal: `\x1b`
- Decimal: `27`

Followed by the command, usually delimited by opening square bracket (`[`) and optionally followed by arguments and the command itself.

Arguments are delimited by semi colon (`;`).

For example:

```sh
\x1b[1;30m  # Set style to bold, black foreground.
\x1b[1;31m  # Set style to bold, red foreground.
\x1b[1;32m  # Set style to bold, green foreground.
\x1b[1;33m  # Set style to bold, yellow foreground.
\x1b[1;34m  # Set style to bold, blue foreground.
\x1b[1;35m  # Set style to bold, magenta foreground.
\x1b[1;36m  # Set style to bold, cyan foreground.
\x1b[1;37m  # Set style to bold, white foreground.
```

## General ASCII Codes

| Name  | decimal | octal | hex  | C-escape | Ctrl-Key | Description                    |
| ----- | ------- | ----- | ---- | -------- | -------- | ------------------------------ |
| `BEL` | 7       | 007   | 0x07 | `\a`     | `^G`     | Terminal bell                  |
| `BS`  | 8       | 010   | 0x08 | `\b`     | `^H`     | Backspace                      |
| `HT`  | 9       | 011   | 0x09 | `\t`     | `^I`     | Horizontal TAB                 |
| `LF`  | 10      | 012   | 0x0A | `\n`     | `^J`     | Linefeed (newline)             |
| `VT`  | 11      | 013   | 0x0B | `\v`     | `^K`     | Vertical TAB                   |
| `FF`  | 12      | 014   | 0x0C | `\f`     | `^L`     | Formfeed (also: New page `NP`) |
| `CR`  | 13      | 015   | 0x0D | `\r`     | `^M`     | Carriage return                |
| `ESC` | 27      | 033   | 0x1B | `<none>` | `^[`     | Escape character               |
| `DEL` | 127     | 177   | 0x7F | `<none>` | `<none>` | Delete character               |

> Note: The **Ctrl-Key** representation is simply associating the non-printable characters
> from ASCII code 1 with the printable (letter) characters from ASCII code 65 ("A").
> ASCII code 1 would be `^A` (Ctrl-A), while ASCII code 7 (BEL) would be `^G` (Ctrl-G).
> This is a common representation (and input method) and historically comes from one of the VT series of terminals.

## Cursor Controls

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[H` | moves cursor to home position (0, 0) |
| `ESC[{line};{column}H` <br> `ESC[{line};{column}f` | moves cursor to line #, column # |
| `ESC[#A` | moves cursor up # lines |
| `ESC[#B` | moves cursor down # lines |
| `ESC[#C` | moves cursor right # columns |
| `ESC[#D` | moves cursor left # columns |
| `ESC[#E` | moves cursor to beginning of next line, # lines down |
| `ESC[#F` | moves cursor to beginning of previous line, # lines down |
| `ESC[#G` | moves cursor to column # |
| `ESC[#;#R` | reports current cursor line and column |
| `ESC[s` | saves the current cursor position |
| `ESC[u` | restores the cursor to the last saved position |

## Erase Functions

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[J` | clears the screen |
| `ESC[0J` | clears from cursor until end of screen |
| `ESC[1J` | clears from cursor to beginning of screen |
| `ESC[2J` | clears entire screen |
| `ESC[K` | clears the current line |
| `ESC[0K` | clears from cursor to end of line |
| `ESC[1K` | clears from cursor to start of line |
| `ESC[2K` | clears entire line |

## Colors / Graphics Mode

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[{...}m` | Set styles and colors for cell and onward. |
| `ESC[0m` | reset all styles and colors |
| `ESC[1m` | set style to bold |
| `ESC[2m` | set style to dim |
| `ESC[2m` | set style to dim |

### Color codes

Most terminals support 8 and 16 colors, as well as 256 (8-bit) colors. These colors are set by the user, but have commonly defined meanings.

#### 8-16 Colors

| Color Name | Foreground Color Code | Background Color Code |
|:------------------|:------------|:------------|
| Black | `30` | `40` |
| Red | `31` | `41` |
| Green | `32` | `42` |
| Yellow | `33` | `43` |
| Blue | `34` | `44` |
| Magenta | `35` | `45` |
| Cyan | `36` | `46` |
| White | `37` | `47` |
| Reset | `0` | `0` |

> Note: The **Reset** is the reset code that resets _all_ colors and text effects.

Most terminals, apart from the basic set of 8 colors, also support the "bright" or "bold" colors. These have their own set of codes, mirroring the normal colors, but with an additional `;1` in their codes:

```sh
# Set style to bold, red foreground.
\x1b[1;31mHello
# Set style to dimmed white foreground with red background.
\x1b[2;37;41mWorld
```

#### 256 Colors

The following escape code tells the terminal to use the given color ID:

```sh
ESC[38;5;${ID}m
```

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[38;5;${ID}m` | Set foreground color. |
| `ESC[48;5;${ID}m` | Set background color. |

Where **ID** is the color index from 0 to 255 of the color table:

![256 Color table](https://user-images.githubusercontent.com/995050/47952855-ecb12480-df75-11e8-89d4-ac26c50e80b9.png)

The table starts with the original 16 colors (0-15).

The proceeding 216 colors (16-231) or formed by a 3bpc RGB value offset by 16, packed into a single value.

The final 24 colors (232-256) are grayscale starting from a shade slighly lighter than black, ranging up to shade slightly darker than white.

Some emulators interpret these steps as linear increments from (`256 / 24`) on all three channels, although some emulators may explicitly define these values.

#### RGB Colors

More modern terminals supports [Truecolor](https://en.wikipedia.org/wiki/Color_depth#True_color_.2824-bit.29) (24-bit RGB), which allows you to set foreground and background colors using RGB.

These escape sequences are usually not well documented.

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[38;2;{r};{g};{b}m` | Set foreground color as RGB. |
| `ESC[48;2;{r};{g};{b}m` | Set background color as RGB. |

> Note that `;38` and `;48` corresponds to the 16 color sequence and is interpreted by the terminal to set the foreground and background color respectively. Where as `;2` and `;5` sets the color format.

## Screen Modes

### Set Mode

| ESC Code Sequence | Description |
|:------------------|:------------|
| `ESC[={value}h` | Changes the screen width or type to the mode specified by value. |
| `ESC[=0h` | 40 x 25 monochrome (text) |
| `ESC[=1h` | 40 x 25 color (text) |
| `ESC[=2h` | 80 x 25 monochrome (text) |
| `ESC[=3h` | 80 x 25 color (text) |
| `ESC[=4h` | 320 x 200 4-color (graphics) |
| `ESC[=5h` | 320 x 200 monochrome (graphics) |
| `ESC[=6h` | 640 x 200 monochrome (graphics) |
| `ESC[=7h` | Enables line wrapping |
| `ESC[=13h` | 320 x 200 color (graphics) |
| `ESC[=14h` | 640 x 200 color (16-color graphics) |
| `ESC[=15h` | 640 x 350 monochrome (2-color graphics) |
| `ESC[=16h` | 640 x 350 color (16-color graphics) |
| `ESC[=17h` | 640 x 480 monochrome (2-color graphics) |
| `ESC[=18h` | 640 x 480 color (16-color graphics) |
| `ESC[=19h` | 320 x 200 color (256-color graphics) |
| `ESC[={value}l` | Resets the mode by using the same values that Set Mode uses, except for 7, which disables line wrapping. The last character in this escape sequence is a lowercase L. |

### Example Modes

| ESC Code Sequence | Description |
| `ESC[?25l` | make cursor invisible |
| `ESC[?25h` | make cursor visible |
| `ESC[?47l` | restore screen |
| `ESC[?47h` | save screen |

### Keyboard Strings

```sh
ESC[{code};{string};{...}p
```

Redefines a keyboard key to a specified string.

The parameters for this escape sequence are defined as follows:

Code is one or more of the values listed in the following table. These values represent keyboard keys and key combinations. When using these values in a command, you must type the semicolons shown in this table in addition to the semicolons required by the escape sequence. The codes in parentheses are not available on some keyboards. ANSI.SYS will not interpret the codes in parentheses for those keyboards unless you specify the /X switch in the DEVICE command for ANSI.SYS.

String is either the ASCII code for a single character or a string contained in quotation marks. For example, both 65 and "A" can be used to represent an uppercase A.

IMPORTANT: Some of the values in the following table are not valid for all computers. Check your computer's documentation for values that are different.

| Key                      | Code     | SHIFT+code | CTRL+code | ALT+code  |
| ------------------------ | -------- | ---------- | --------- | --------- |
| F1                       | 0;59     | 0;84       | 0;94      | 0;104     |
| F2                       | 0;60     | 0;85       | 0;95      | 0;105     |
| F3                       | 0;61     | 0;86       | 0;96      | 0;106     |
| F4                       | 0;62     | 0;87       | 0;97      | 0;107     |
| F5                       | 0;63     | 0;88       | 0;98      | 0;108     |
| F6                       | 0;64     | 0;89       | 0;99      | 0;109     |
| F7                       | 0;65     | 0;90       | 0;100     | 0;110     |
| F8                       | 0;66     | 0;91       | 0;101     | 0;111     |
| F9                       | 0;67     | 0;92       | 0;102     | 0;112     |
| F10                      | 0;68     | 0;93       | 0;103     | 0;113     |
| F11                      | 0;133    | 0;135      | 0;137     | 0;139     |
| F12                      | 0;134    | 0;136      | 0;138     | 0;140     |
| HOME (num keypad)        | 0;71     | 55         | 0;119     | \--       |
| UP ARROW (num keypad)    | 0;72     | 56         | (0;141)   | \--       |
| PAGE UP (num keypad)     | 0;73     | 57         | 0;132     | \--       |
| LEFT ARROW (num keypad)  | 0;75     | 52         | 0;115     | \--       |
| RIGHT ARROW (num keypad) | 0;77     | 54         | 0;116     | \--       |
| END (num keypad)         | 0;79     | 49         | 0;117     | \--       |
| DOWN ARROW (num keypad)  | 0;80     | 50         | (0;145)   | \--       |
| PAGE DOWN (num keypad)   | 0;81     | 51         | 0;118     | \--       |
| INSERT (num keypad)      | 0;82     | 48         | (0;146)   | \--       |
| DELETE (num keypad)      | 0;83     | 46         | (0;147)   | \--       |
| HOME                     | (224;71) | (224;71)   | (224;119) | (224;151) |
| UP ARROW                 | (224;72) | (224;72)   | (224;141) | (224;152) |
| PAGE UP                  | (224;73) | (224;73)   | (224;132) | (224;153) |
| LEFT ARROW               | (224;75) | (224;75)   | (224;115) | (224;155) |
| RIGHT ARROW              | (224;77) | (224;77)   | (224;116) | (224;157) |
| END                      | (224;79) | (224;79)   | (224;117) | (224;159) |
| DOWN ARROW               | (224;80) | (224;80)   | (224;145) | (224;154) |
| PAGE DOWN                | (224;81) | (224;81)   | (224;118) | (224;161) |
| INSERT                   | (224;82) | (224;82)   | (224;146) | (224;162) |
| DELETE                   | (224;83) | (224;83)   | (224;147) | (224;163) |
| PRINT SCREEN             | \--      | \--        | 0;114     | \--       |
| PAUSE/BREAK              | \--      | \--        | 0;0       | \--       |
| BACKSPACE                | 8        | 8          | 127       | (0)       |
| ENTER                    | 13       | \--        | 10        | (0        |
| TAB                      | 9        | 0;15       | (0;148)   | (0;165)   |
| NULL                     | 0;3      | \--        | \--       | \--       |
| A                        | 97       | 65         | 1         | 0;30      |
| B                        | 98       | 66         | 2         | 0;48      |
| C                        | 99       | 66         | 3         | 0;46      |
| D                        | 100      | 68         | 4         | 0;32      |
| E                        | 101      | 69         | 5         | 0;18      |
| F                        | 102      | 70         | 6         | 0;33      |
| G                        | 103      | 71         | 7         | 0;34      |
| H                        | 104      | 72         | 8         | 0;35      |
| I                        | 105      | 73         | 9         | 0;23      |
| J                        | 106      | 74         | 10        | 0;36      |
| K                        | 107      | 75         | 11        | 0;37      |
| L                        | 108      | 76         | 12        | 0;38      |
| M                        | 109      | 77         | 13        | 0;50      |
| N                        | 110      | 78         | 14        | 0;49      |
| O                        | 111      | 79         | 15        | 0;24      |
| P                        | 112      | 80         | 16        | 0;25      |
| Q                        | 113      | 81         | 17        | 0;16      |
| R                        | 114      | 82         | 18        | 0;19      |
| S                        | 115      | 83         | 19        | 0;31      |
| T                        | 116      | 84         | 20        | 0;20      |
| U                        | 117      | 85         | 21        | 0;22      |
| V                        | 118      | 86         | 22        | 0;47      |
| W                        | 119      | 87         | 23        | 0;17      |
| X                        | 120      | 88         | 24        | 0;45      |
| Y                        | 121      | 89         | 25        | 0;21      |
| Z                        | 122      | 90         | 26        | 0;44      |
| 1                        | 49       | 33         | \--       | 0;120     |
| 2                        | 50       | 64         | 0         | 0;121     |
| 3                        | 51       | 35         | \--       | 0;122     |
| 4                        | 52       | 36         | \--       | 0;123     |
| 5                        | 53       | 37         | \--       | 0;124     |
| 6                        | 54       | 94         | 30        | 0;125     |
| 7                        | 55       | 38         | \--       | 0;126     |
| 8                        | 56       | 42         | \--       | 0;126     |
| 9                        | 57       | 40         | \--       | 0;127     |
| 0                        | 48       | 41         | \--       | 0;129     |
| \-                       | 45       | 95         | 31        | 0;130     |
| \=                       | 61       | 43         | \---      | 0;131     |
| \[                       | 91       | 123        | 27        | 0;26      |
| \]                       | 93       | 125        | 29        | 0;27      |
|                          | 92       | 124        | 28        | 0;43      |
| ;                        | 59       | 58         | \--       | 0;39      |
| '                        | 39       | 34         | \--       | 0;40      |
| ,                        | 44       | 60         | \--       | 0;51      |
| .                        | 46       | 62         | \--       | 0;52      |
| /                        | 47       | 63         | \--       | 0;53      |
| \`                       | 96       | 126        | \--       | (0;41)    |
| ENTER (keypad)           | 13       | \--        | 10        | (0;166)   |
| / (keypad)               | 47       | 47         | (0;142)   | (0;74)    |
| \* (keypad)              | 42       | (0;144)    | (0;78)    | \--       |
| \- (keypad)              | 45       | 45         | (0;149)   | (0;164)   |
| \+ (keypad)              | 43       | 43         | (0;150)   | (0;55)    |
| 5 (keypad)               | (0;76)   | 53         | (0;143)   | \--       |

## Resources

- [logrusorgru/aurora Golang ultimate ANSI-colors that supports Printf/Sprintf methods](https://github.com/logrusorgru/aurora)
- [Wikipedia: ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code)
- [Build your own Command Line with ANSI escape codes](http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html)
- [ascii-table: ANSI Escape sequences](http://ascii-table.com/ansi-escape-sequences.php)
- [bluesock: ansi codes](https://bluesock.org/~willkg/dev/ansi.html)
- [bash-hackers: Terminal Codes (ANSI/VT100) introduction](http://wiki.bash-hackers.org/scripting/terminalcodes)

## Colors（几乎所有颜色的中英文对照）

https://bbs.fobshanghai.com/thread-3632955-1-1.html?btwaf=61056113

amber 琥珀色(黄色)
amethyst 紫(水晶)色 antique brass 青古铜色 antique golden 古铜色 antique violet 古紫色 antique white 古董白 apricot 杏黄 aqua green 水绿色,浅绿色
aquamarine 碧绿色 aquamarine blue 蓝绿色 auburn 赤褐色 august green 深绿色 autumn mink 深咖啡色 azure 天蓝色 azure green 碧绿色 baby blue 淡蓝色
baby pink 浅粉红色 bark 树皮色 begin colour 自然色 beige 浅褐色,米白色,灰褐色 benzo blue 靛青色 berry 鲜红色,浆果红 biscuit 淡褐色 bisque 桔黄色 black 黑色
blue 蓝色 blue green 竹青色,青绿色 blueviolet 紫罗兰色 bluish 带蓝色的，浅蓝色 blush 淡红色 bone 骨色 bottle green 深绿色 brick red 青莲色 bronze 青铜色
brown 褐色,棕色,茶色 buff 浅黄色;暗黄色 burgundy 葡萄红;枣红色 burly wood 实木色 butter 油黄色 butter cup 深黄色 cadet blue 军蓝色,灰蓝色 calamineblue
淡蓝色 camel 驼色 camouflage 迷彩色 caramel 酱色 carmine 深红色,洋红色 carnation 粉红色,康乃馨 celery 菜绿色,芹菜 celeste 天蓝色 cerise 樱桃色 chalky 白垩的
chambray 有条纹或格子花纹的布 charcoal 炭色 charcoal gray 炭灰色 chartreuse 黄绿色 cherry 鲜红色,樱桃色 chestnut 栗色 chocolate 红褐色,赭石色,巧克力色 chost
white 幽灵白 cinnamon 肉桂色 citrine 柠檬黄 citrus green 浅绿色 classic navy 蓝色 clay 泥土 clear 透明 cobalt 钴蓝色,深蓝色 cobalt blue 钴蓝色,艳蓝色
cochineal 胭脂红,洋红 cocoa 可可色,黄棕色 coffee 咖啡色 cold colour 冷色 colour combination 配色 colour matching 配色 colour mixing 调色
complementary colour 补色 contracting colour 收缩色 contrast color 衬色 copper 铜;红铜色 copper red 铜红色 coral 珊蝴色 coral haze 桔红色
cornflower blue 菊蓝色;浅蓝色 cornsilk 米绸色 cream 奶油色，米黄色,淡黄色 cream soda 棕色 crepe 透明 crimson 深红色 crystal cream 奶油白 crystaline
水晶色 cyan 青色,蓝绿色 daffod;daffadilly 水仙,鲜黄色 dark 深色 deep 深色 delicate color 娇色 denim 粗斜纹棉布,蓝牛仔布 dim gray 暗灰色 dodger blue 闪蓝色
dove 鸽子;乳白色 drab 土褐色 dry rose 浅紫色 dull silver 雾银 dun 焦茶色,暗褐色 ecru 本色的，淡褐色,米色 emerald 祖母绿 emerald green 鲜绿色,艳绿色 espresso
浓咖啡色 essential colour 基本色 expansive colour 膨胀色 fashion colour 流行色 firebrick 火砖色 flake white 铅白,片白 flax 亚麻布,淡黄色
fleshcolor 肉色 floral 花的，花似的 floral white 花白色 florid 鲜红色 forest green 森林绿 fuchsia 紫红色，粉玫色 full colour 彩色 fundamental
colour 原色 fuscous 暗褐色,深色 garnet 深红色 geranium 原色红；天竺葵 gilt 青铜色 global brown 咖啡色 golden 金色,金黄色 goldenrod 金麒麟 色 gray;grey
灰色,灰白 green 绿色 greenyellow 黄绿色 ground colour 底色 gunmetal 青铜色，黑镍色 havana 雪茄色 hazel 赤褐色 hepatic 猪肝色 honeydew 蜜色 hot pink
艳粉色,粉红色 hyacinth 紫蓝色 ice 冰色 iced coffee 冰咖色 incarnadine 肉色,粉红色 indigo 靛青色 inter colour 国际流行色 intermedium colour 中间色
ivory 乳白色，象牙色 jade 碧玉色，翡翠 色 kelly 黄绿色，鲜绿色 khaki 黄褐色，卡其色 lavender 淡紫色,藕色 lavender blush 淡紫红 lawn green 草绿色 lemon 柠檬色
lemon chiffon 柠檬绸色 lias 淡紫色 light 浅的 light cyan 浅青色 lilac 淡紫色 lily blush 嫩粉色 lily orange 桔色 lily sky 粉蓝色 lime 淡黄绿色 lime
green 橙绿色 linen 亚麻色 loden 深橄榄色;深绿色 lyons blue 蓝紫色 magenta 洋红 magenta 红紫色,洋红 maroon 栗色,褐红色 mauve 淡紫色,紫红 melon (各种的)瓜
metal colour 金属色 mint cream 薄荷色 misty gray 雾灰色 mistyrose 浅玫瑰色 mixtz 混色 moccasin 鹿皮色 modena 深紫色 moss green 苔绿色
multicolour 多种色彩 mustard 深黄色 natural 自然色 navajo white 纳瓦白 navy 深蓝色 navy blue 深蓝色,藏青色;海蓝 neutral colour 中间色 nickel 镍色
oatmeal 灰白色 off white 灰白色 oldlace 老白色 olive 橄榄绿 olive drab 草绿色;深绿褐色 olive green 橄榄绿，茶绿色 opaque 不透明 orange 橙色,桔黄色 orchid
淡紫色 oriental red 大红 oxblood 无光泽的深红色 oyster grey 米灰色 oyster white 乳白色 pale 淡的,苍白色 pale gold 金色 pale green 苍绿色 pale hay
干草色 pale oink kiss 浅粉红色 pale turquoise 苍绿色 pale violet red 苍紫罗兰色 palegoldenrod 苍麒麟色 pansy 紫罗兰色 pea green 淡绿色,青豆色 peach
桃红色 peacock blue 孔雀蓝 pearl grey 珠光灰 pearlied gold 珠光金色 periwinkle 花布 pewter 粉红色 ：pink 红色 Red 紫色 Purple 蓝色 Blue 玫红 Rose
粉红 Pink 彩叶草 Colens 标准混色 Formula Mixed 古铜色 Bronze 柠檬黄 Lemon 魔力 Magic 摩西 Mosaic

桔黄色 Orange 天鹅绒红 Red Velvet 玫瑰红 Rose 红宝石红 Ruby

橙红玫瑰红 Salmon Rose

黄色 Yellow “绚丽”彩虹系列 “Superfine Rainbow” 橙红色带花边 Salmon Lace

华丽 Color Pride 欢乐之舞 Festive Dance 七彩彩虹 Multicolor Rainbow 红天鹅绒 Red Velvet 混色 Mixed 白色 White 深粉红色 Blush pink 洋红色 Carmine
粉红色 Pink 粉红色眼 Pink Eye 桔红色 Orange 黄色 Yellow 黄色黑心 Yellow With Black Center 橙色黑心 Orange With Black Center 猩红Scarlet 紫色
Purple 橙红 Salmon 白色红心 White with Eye 大红 Red 橙红 Salmon 粉色红眼 Shell Pink 紫色溅开色 Purple Splash 粉红溅开色 Pink Splash 橙红溅开色 Salmon
Splash 深红色 Crimson 红色白边 Picotee 紫红色 Purple 桃红色 Carmine Rose 紫红色白边 Purple Picotee 淡腥红色 Pastel Scarlet 桔黄色渐变色 Orange
Shades 玫瑰色渐变色 Rose Shades 玫瑰与粉红色 Rose and Pink

蓝与白色 Blue and White 红与白色 Red and White 淡蓝色 Light Blue 天蓝渐变色 Sky Blue Shades 标准混色 Formula Mixed 纯蓝 True Blue 纯蓝带白心 True
Blue With White Throat 红色带白 Red With White Edge 玫瑰红带白心 Rose Red With White Throat 古代稀 Godetia

深玫瑰红 Deep Rose 淡紫色 Lavender

淡玫瑰色 Lilac Rose 橙红色 Salmon 淡紫色 Lavender 肉色 Shell Pink 羞红色 Bulsh 洋红色玫瑰色 Carmine Rose 珊瑚红 Coral 深橙色 Deep Salmon 淡紫色
Lavender 浅粉红色 Light Pink 霓红玫瑰红 Neon Rose 桔红色 Orange 蓝色 Blue 粉红色 Pink 白色 White 黄色 Yellow 蓝色 Blue 蓝色镶边 Blue Picotee 淡蓝色
Light Blue 纯白色 Pure White 粉红色镶边 Pink Picotee 淡玫瑰红 Lilac Rose 单边多枝性“海蒂系列” 标准混色 Formula Mixed 橙红色 Salmon 蓝色镶边 Blue Rim lmp
樱花红 Cherry Blossom 深蓝色 Deep Blue 天蓝色 Sky Blue 淡紫色 Orchid 淡蓝色 Pastel Blue 粉红色镶边 Pink Rim 淡玫瑰色 Lilac Rose 蓝色镶边 Blue Rim
淡蓝色 Light Blue 玫瑰镶边 Rose Rim 亮黄色 Spry 焰红色 Flame 纯黄 Yellow 蓝帽子 Blue Cap 淡黄色 Canary 焰红色 Flame 海蓝 Ocean 红帽子 Red Cap 雪白 Snow
蓝色渐变色 Blue Shades 玫瑰渐变色 Rose Shades 红黄双色 Red and Yellow Bicolor 猩红色渐变色 Scarlet Shades 白色带斑 White with Blotch 黄色带斑 Yellow
with Blotch 桔黄色带斑 Orange with Blotch 紫色/黄色带斑 Purple/Yellow with Blotch 白色带玫瑰色斑 White with Rose Blotch 蓝色带斑 Blue with
Blotch 玫瑰色带斑 Rose with Blotch 猩红色带斑 Scarlet with Blotch 深蓝色带斑 Deep Blue with Bblotch 红色/黄色带斑 Red/Yellow with Blotch 白色带斑
White with Blotch 黄色带斑 Yellow with Blotch 黄色带红斑 Yellow with Red Blotch 天蓝色 Azure 奶油色 Cream 金色 Golden 海蓝色 Marina 蓝色黄色
Blue & Yellow 丝绸色 Chiffon 深蓝色 Deep Blue 红色黄色 Red & Yellow 奶白色 Sherbet 霞红色 Sunset 黄色 Yellow 改良深蓝 Deep Blue Imp New 改良纯蓝
True Blue New 蓝芯 Blue Center 桔黄色 Orange 樱草色 Primrose 天蓝色 Sky Blue 粉红色 Blush Pink 红葡萄酒色 Burgundy 珊瑚色 Coral 淡紫色 Lilac 深玫瑰色
Deep Rose 淡橙红色 Pastel Salmon 晨粉红 Pink Morn 晨红色 Red Morn 叶脉标准混色 Vein Formula Mixed

一. 红色类

红色 red 朱红 vermeil; vermilion; ponceau 粉红 pink; soft red; rose bloom 梅红 plum;crimson;fuchsia red 玫瑰红 rose madder; rose 桃红
peach blossom; peach; carmine rose 樱桃红 cherry; cerise 桔红 reddish orange; tangerine; jacinth; salmon pink; salmon 石榴红
garnet 枣红 purplish red; jujube red; date red 莲红 lotus red 浅莲红 fuchsia pink 豉豆红 bean red 辣椒红 capsicum red 高粱红 Kaoliang
red 芙蓉红 hibiscus red; poppy red; poppy 胭脂红 rogue red carmine; cochineal; lake 鲑鱼红 salmon 玳瑁红 hawksbill turtle red 海螺红
cadmium orange 宝石红 ruby red 玛瑙红 agate red 珊瑚红 coral 金红 bronze red 铁红 iron oxide red 铁锈红 rust red 镉红 cadmium red 铬红
chrome red 砖红 brick red 土红 laterite; reddle 郎窑红 lang-kiln red 均红 Jun-kiln red 釉底红 underglaze red 威尼斯红 Venetian red 法国红
French vermilion 茜红 alizarin red; madder red 洋红 carmine; magenta 品红 pinkish red; magenta 猩红 scarlet red; scarlet; blood
red 油红 oil red 紫红 purplish red; madder red; wine red; wine; carmine; amaranth; claret; fuchsia; magenta; heliotrope;
mauve 玫瑰紫红 rose carmine; rose mauve 深紫红 prune; mulberry 深藕红 conch shell 棕红 henna 暗红 dark red; dull red 鲜红 scarlet red;
scarlet; bright red; fresh red; blood red; madder; ruby; cerise; cherry 血红 blood red; incarnadine 血牙红 shell pink; peach
beige 绯红 scarlet; crimson; geranium pink 米红 silver pink 红 deep red; crimson 淡红 light red; carnation

二.橙色类

橙色 orange 三.黄色类 黄色 yellow 桔黄 orange; crocus; gamboge; cadmium orange 深桔黄,深橙 deep orange 浅桔黄,浅橙 clear orange; light
orange; rattan 柠檬黄 lemon yellow lemon citrine citron 玉米黄 maize 榄黄 olive yellow 樱草黄 primrose yellow 稻草黄 straw yellow 芥末黄
mustard 杏黄 apricot; apricot buff; bronze yellow 蛋黄 vitelline; yolk yellow;egg yellow 藤黄 rattan yellow 鳝鱼黄 eel yellow 牙黄
ivory 日光黄 sunny yellow 石黄 mineral yellow 土黄 earth yellow; yellowish brown; yellow ocher; golden apricot 砂黄 sand yellow
金黄 golden yellow, gold 铁黄 iron oxide yellow; iron buff 镉黄 cadmium yellow 铬黄 chrome yellow 黄 cobalt yellow 深黄,暗黄 deep
yellow 棕黄 tan 青黄 bluish yellow 黄 isabel sallow grey yellow 米黄 apricot cream cream 嫩黄 yellow cream 鲜黄 cadmium yellow
canary 黄 light yellow 中黄 midium yellow 浅黄 light yellow;pale yellow;buff 淡黄 jasmin(e); primrose

四.绿色类 绿色 green 豆绿 pea green bean green 豆绿 light bean green; asparagus green 橄榄绿 olive green olive 茶绿 tea green celandine
green plantation 葱绿 onion green pale green 苹果绿 apple green 野绿 field green 林绿 forest green 洋蓟绿 artichoke green 苔藓绿 moss
green bracken green 草地绿,草绿 grass green meadow green oliver green olive drab 水草绿 water grass green 深草绿 jungle green 灰湖绿
agate green 水绿 aqua green 海水绿 marine green 酸性绿 acid green 水晶绿 crystal green 玉绿 jade green 石绿 mineral green 松石绿
spearmint; viridis 铜绿 verdigris 铜锈绿 patina green 镉绿 cadmium green 铬绿 chrome green 钴绿 cobalt green 孔雀绿 peacock green 威尼斯绿
Venetian green 巴黎绿 Paris green king's green 墨绿 blackish green green black; jasper; dark green; deep green 墨玉绿 emerald
black 深绿 dark green petrol; Chinese green; bottle green 暗绿 sap green dark green deep green 青绿 dark green 碧绿 azure green;
turquoise green viridity 翠绿 emerald green; jade green bright green verdancy viridity 深翠绿 viridian 蓝绿 blue green
aquamarine 黄绿 yellow green 灰绿 grey green sage green hedge green; mignonette; sea spray; celadon 褐绿 breen 品绿 light green
malachite green 鲜绿 clear green; emerald green vivid green 嫩绿 pomona green verdancy 中绿 medium green; golf green 绿 light
green 淡绿 pale green

五.青色类 青色 cerulean blue blue green 青 pea green; bean green 花青 flower blue 茶青 tea green 葱青 onion green 天青 celeste; azure
霁青 sky-clearing blue 石青 mineral blue 铁青 electric blue river blue 蟹青 turquoise ink blue 鳝鱼青 eel green 青 egg blue 影青 misty
blue; white blue 黛青 bluish 群青,伟青 ultramarine 暗青 dark blue; deep cerulean 青 navy blue; dark blue; Ming blue 靛青 indigo 大青
smalt 粉青 light greenish blue 鲜青 clear cerulean 青 light blue; light cerulean 淡青 pale cerulean light greenish blue

六.蓝色类 蓝色 blue 天蓝 sky blue; azure celeste; azure cerulean blue; arisian blue 蔚蓝 azure; sky blue 月光蓝 moon blue 海洋蓝 ocean
blue 海蓝 sea blue 蓝 acid blue 深湖蓝 vivid blue 中湖蓝 bright blue 浅湖蓝 canal blue 清水蓝 water blue

雪蓝 ice-snow blue 孔雀蓝 peacock blue 宝石蓝 sapphire; jewelry 末蓝 powder blue 铁蓝 iron blue 钴蓝 cobalt blue king's blue 普鲁士蓝
Prussian blue 北京蓝 Beijing blue 林蓝 indanthrene blue 品蓝 reddish blue royal blue;king's blue 靛蓝 indigo; indigo blue; benzo
blue 菘蓝 woaded blue 磨蓝 stone-washed indigo 藏蓝 purplish blue; navy blue; navy 海军蓝 navy blue; navy 宝蓝 royal blue 墨蓝 blue
black 蓝 turquoise blue 紫蓝 hyacinth;purplish blue 浅紫蓝 Dutch blue 青蓝 ultramarine 深灰蓝 blue ashes 深蓝 deep blue; dark blue
navy blue mandarin blue Antwerp blue mazarine smalt ultramarine 暗蓝 deep blue; dark blue 鲜蓝 clear blue 中蓝 medium blue
azure blue 浅蓝 light blue 淡蓝 pale blue baby blue calamine blue

七.紫色类 紫色 purple; violet 紫罗兰色 violet 紫藤色 lilac 紫水晶色 amethyst 葡萄紫 grape 茄皮紫 aubergine; wineberry 玫瑰紫 rose violet 丁香紫 lilac
钴紫 cobalt violet 墨紫 violet black 绛紫 dark reddish purple 暗紫 violet deep; dull purple; damson 乌紫 raisin 蓝紫 royal light 鲜紫
violet light 深紫 amaranth; modena 浅紫 grey violet 淡紫 pale purple lavender; lilac; orchid 淡白紫 violet ash 青莲 pale purple;
heliotrope 青莲 amaranth purple 雪青 lilac 墨绛红 purple black 暗绛红 purple deep 浅绛红 purple light

八.黑色类 色 black 黑 carbon black;charcoal black 暗黑 pitch-black ; pitch-dark 漆黑dull black 白色 white 象牙白 ivory white 牡蛎白 oyster
white 珍珠白 pearl white 玉石白 jade white 银白 silver white 羊毛白 wool white 乳白 milky white 米白off-white; shell 雪白 snow-white 灰白
greyish white 纯白 pure white 本白 rawwhite ;off white 粉红白 pinky white 浅紫白 lilac white 灰色 grey 银灰色 silvergrey 炭灰色 charcoal
grey 烟灰 smoky grey 雾灰 misty grey 黑灰 grey black 金色gold 银色 silver 青古铜色 bronze;bronzy 驼色 camel ;light tan 米色
beige;cream;gray sand 卡其色 khaki 豆沙色 cameo brown 水晶色 crystal 荧光色iridescent 茶褐 umber;auburn 淡褐 light brown
