# make compiled languages more oop

## Python Result (Expected)

```shell
foo from animal
Pruttel
foo from duck
Pruttel
```

With the `Duck` class only overrides `foo` methods, the overridden `foo` methods is used when `Duck().sound()` is called.

Thins are quite different for CPP and Go.

## C++ Result

```shell
foo from animal
Pruttel
foo from animal
Pruttel
```

## Go Result

```shell
foo from animal
Pruttel
foo from animal
Pruttel
```

