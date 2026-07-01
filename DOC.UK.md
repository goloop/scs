# scs — довідник

Повний довідник пакета `scs`: ментальна модель, токенізатор, усі стилі й
конвертери, `Caser` та ініціалізми, детекція, гарантії й обмеження, практичні
рецепти.

Англійська версія: **[DOC.md](DOC.md)**.

## Зміст

- [Ментальна модель](#ментальна-модель)
- [Токенізатор: Split і Words](#токенізатор-split-і-words)
- [Стилі й конвертери](#стилі-й-конвертери)
- [Вибір стилю під час виконання](#вибір-стилю-під-час-виконання)
- [Ініціалізми й Caser](#ініціалізми-й-caser)
- [Числа](#числа)
- [Детекція й предикати](#детекція-й-предикати)
- [Гарантії й обмеження](#гарантії-й-обмеження)
- [Рецепти й поради](#рецепти-й-поради)

## Ментальна модель

`scs` конвертує ідентифікатори між конвенціями іменування: `camelCase`,
`PascalCase`, `snake_case`, `kebab-case`, `SCREAMING_SNAKE_CASE`, `dot.case`,
`Title Case` і `Sentence case`.

Увесь пакет спирається на одну ідею: **є єдиний універсальний токенізатор.**
`Split` зводить будь-який ввід до списку нормалізованих слів, а кожен стиль — це
просто інший рендеринг цього списку. Оскільки один токенізатор живить кожен
рендерер:

- Конвертери **тотальні** — вони ніколи не повертають помилку й ніколи не мусять
  знати початковий стиль вводу.
- Додати стиль означає додати рендерер, а не ще один попарний конвертер N×N.

```go
scs.ToSnake("HTTPServerID") // "http_server_id"
scs.ToCamel("user_id")      // "userId"
scs.ToKebab("HelloWorld")   // "hello-world"
```

```go
import "github.com/goloop/scs/v2"
```

## Токенізатор: Split і Words

```go
func Split(s string) []string
func Words(s string) iter.Seq[string]
```

`Split` повертає нормалізовані слова; `Words` видає їх ліниво як
`iter.Seq[string]`:

```go
scs.Split("HTTPServerID")   // ["http", "server", "id"]
scs.Split("user_id")        // ["user", "id"]
scs.Split("web2print")      // ["web2print"]

for w := range scs.Words("parseJSONResponse") {
    fmt.Println(w) // parse, json, response
}
```

Межі слів розставляються:

- на роздільниках (`_`, `-`, `.`, пробіл);
- на переходах нижній→верхній регістр (`helloWorld` → `hello|World`);
- у кінці акроніма перед словом з малих літер (`HTTPServer` → `HTTP|Server`);
- перед Title-словом, що починається з цифри (`v2Final` → `v2|Final`).

## Стилі й конвертери

Кожен стиль має тотальну функцію `To…`:

| Функція | Приклад виводу |
|---------|----------------|
| `ToCamel`          | `helloWorld` |
| `ToPascal`         | `HelloWorld` |
| `ToSnake`          | `hello_world` |
| `ToKebab`          | `hello-world` |
| `ToScreamingSnake` | `HELLO_WORLD` |
| `ToDot`            | `hello.world` |
| `ToTitle`          | `Hello World` |
| `ToSentence`       | `Hello world` |

```go
scs.ToCamel("hello_world")      // helloWorld
scs.ToPascal("hello-world")     // HelloWorld
scs.ToScreamingSnake("userID")  // USER_ID
scs.ToSentence("hello_world")   // Hello world
```

## Вибір стилю під час виконання

Коли цільовий стиль надходить із конфігу, CLI-прапора чи бази, використовуйте
енум `Style` і `Convert`:

```go
func Convert(to Style, s string) string
func ParseStyle(name string) (Style, bool)
func (s Style) String() string
func (s Style) Valid() bool
```

```go
style, ok := scs.ParseStyle("kebab")     // Kebab, true
out := scs.Convert(style, "HTTPServerID") // "http-server-id"
```

Константи `Style`: `Unknown`, `Camel`, `Pascal`, `Snake`, `Kebab`,
`ScreamingSnake`, `Dot`, `Title` і `Sentence`. `Unknown` — це нульове значення й
відповідь «немає єдиного стилю» від `Detect`.

## Ініціалізми й Caser

За замовчуванням слова в Title-регістрі (`Id`, `Url`, `Http`), що завжди
повертається туди-назад. Щоб дотримуватися конвенції Go з великих ініціалізмів,
збудуйте `Caser`:

```go
func New(opts ...Option) *Caser
func WithAcronyms(words ...string) Option
```

```go
c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP", "API"))

c.ToPascal("user_id")         // "UserID"
c.ToCamel("http_url_builder") // "httpURLBuilder" (перше слово лишається малим)
```

`Caser` надає ті самі методи `To…` і `Convert`, що й пакетні функції. Він
**незмінний і безпечний для конкурентного використання** — збудуйте раз і
діліться.

Ініціалізми — опційні, бо суміжні акроніми з великих літер не завжди можна знову
розділити (`"HTTPAPI"` неоднозначне), тож безпечний для round-trip Title-регістр
є дефолтом.

## Числа

Цифри чіпляються до сусідніх літер і ніколи не ділять слово самі по собі, тож
фрагменти ідентифікаторів лишаються цілими:

```go
scs.ToSnake("web2print") // "web2print"
scs.ToSnake("sha256sum") // "sha256sum"
scs.ToSnake("oauth2")    // "oauth2"
```

Лише явний роздільник перетворює число на окреме слово:

```go
scs.ToSnake("web 2 print") // "web_2_print"
```

## Детекція й предикати

```go
func Detect(s string) (Style, bool)
func Is(style Style, s string) bool
func IsCamel(s string) bool // IsSnake, IsKebab, IsPascal, IsScreamingSnake,
                            // IsDot, IsTitle, IsSentence
```

`Detect` повертає стиль **лише коли ввід канонічний рівно для одного стилю**.
Неоднозначні входи (як голе слово `"api"`, валідне в кількох стилях одразу)
повертають `(Unknown, false)`:

```go
scs.Detect("user_id") // Snake, true
scs.Detect("userId")  // Camel, true
scs.Detect("USER_ID") // ScreamingSnake, true
scs.Detect("api")     // Unknown, false
```

Постильові предикати (і загальний `Is(style, s)`) повідомляють, чи рядок уже
канонічний для цього стилю — корисно, щоб пропустити зайву конвертацію.

## Гарантії й обмеження

- Конвертери **ніколи не панікують** і завжди повертають валідний UTF-8.
- `snake_case`, `kebab-case` і `dot.case` **ідемпотентні** й зберігають точну
  послідовність слів для будь-якого вводу, тож прогін значення крізь них
  безвтратний.
- Для звичайних багатолітерних ідентифікаторів кожен стиль повертається
  туди-назад (`snake → camel → snake` дає оригінал).
- Регістрові злиті стилі (`camelCase`, `PascalCase`, `Title Case`) не можуть бути
  ідемпотентними для *однолітерних* чи суміжних акронімів з великих літер, бо
  такі рендеринги за природою неоднозначні з акронімами — це властивість
  будь-якого конвертера без словника, а не дефект.
- Токенізатор без словника не відрізнить `IPv6` від патерну «акронім + слово з
  малих літер», тож `ToSnake("IPv6Address")` дає `i_pv6_address`. Використовуйте
  `WithAcronyms` чи роздільник, коли це важливо.

## Рецепти й поради

**Спершу нормалізуйте до безвтратного стилю.** Коли потрібен стабільний ключ,
конвертуйте в `snake_case`/`kebab-case`/`dot.case` — вони ідемпотентні й
зберігають послідовність слів, тож значення переживає повторну обробку.

**Діліться одним Caser.** Збудуйте `scs.New(scs.WithAcronyms(...))` раз на старті
й повторно використовуйте всюди; він незмінний і потокобезпечний.

**Кермуйте стилем із даних.** Зберігайте чи отримуйте ім'я стилю, розв'язуйте
його через `ParseStyle` і рендеріть через `Convert` — без switch по восьми
функціях.

**Пропускайте зайві конвертації.** Захищайтеся предикатом `Is…` (чи
`Is(style, s)`), коли хочете конвертувати лише не-канонічні значення.

**Розв'язуйте неоднозначні техтерміни.** Для токенів на кшталт `IPv6` чи `OAuth2`
подайте роздільник у джерелі або зареєструйте акронім через `WithAcronyms`.
