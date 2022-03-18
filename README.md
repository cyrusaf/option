# Option

## Usage

```golang
func (u UserStore) GetUser(id string) Option[User] {
    if user, ok := u.data[id]; ok {
        return option.Some(user)
    }
    return option.None[User]()
}
...

user, ok := u.GetUser("1").Unwrap()
if !ok {
    return errors.New("user not found")
}
fmt.Printf("username: %s\n", user.Username)
```

## Why use the Option type?

It is common to use a pointer to denote an optional type in idiomatic Go today:

```golang
func GetUser(id string) *User { ... }
...
user := GetUser("1")
if PasswordMatches(user.HashedPassword, password) { ... }
```

If the above example, if the `GetUser` function returned `nil`, we would have a
runtime panic in our code. If you were a reviewer of the last two lines of code,
it isn't clear that `GetUser` returns an optional type and it would be easy to
miss making the comment "you should handle the case of `GetUser` returning nil."

Rust solves this issue using the `Option<T>` type alongside [Pattern Matching](https://doc.rust-lang.org/book/ch06-02-match.html#matching-with-optiont) to enforce that blocks of code with access to the value of `T` only get run if the `Option<T>` is `Some`.

```rust
fn get_user(id: String) -> Option<User> { ... }
...
let user = get_user("1")
if let Some(user) = user {
    if password_matches(user.hashed_password, password) { ... }
}
```

Go does not have pattern matching, so we cannot add compiler enforcement (without going too far into functional territory). But, enforcement can be handled by static analysis tools and we can still solve the readability issue in an idiomatic way:

```golang
func GetUser(id string) Option[User] { ... }
...
userOption := GetUser("1")
user, ok := userOption.Unwrap()
if !ok {
    // return 404
}
if PasswordMatches(user.HashedPassword, password) { ... }
```

The code above forces the `.Unwrap()` call to get access to the value
of the user and forces both the writer and reader to think about how to handle `!ok`. It becomes obvious to the code reviewer if the writer skips handling the `None` case: `user, _ := userOption.Unwrap()`.

### Alternatives

The function below that returns a `(User, bool)` achieves a similar goal without
the extra `Option[T]` type.

```golang
func GetUser(id string) (User, bool) { ... }
```

But `Option[T]` feels better for two reasons:

1. It can be used as a field in a struct or a function argument:

```golang
type Response struct {
    User Option[T]
}

func New(region Option[string]) { ... }
```

2. `Option[T]` is self documenting. It avoids needing the comment `// return (User, true) if a user is found or (User, false) if a user is not found`


## Why no Result type?

Errors in Golang already achieve the readability goal. If there is an error,
it is clear to both the writer and reader that there is an error and it needs to
be handled:

```golang
user, err := GetUser("1")
if err != nil { ... }
```
