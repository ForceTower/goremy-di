# Changelog

## 20220524 - remy/v1.0.0

- Core release
- Register and Retrieve using `generics`
- Example in repository

## 20220530 - remy/v1.1.0

- Create README with more detailed instructions
- Fix some typos
- Add method `SetGlobalInjector`
- Allow to define a **ParentInjector** in `Injector` constructor
- Add more documentation in public types/methods
- Create methods `GetGen` and `GetGenFunc` to pass values dynamically

## 20220721 - remy/v1.2.0

- Replace default type resolution - Now it will not use the `reflect` _package by default_.
- Add `UseReflectionType` option in `Config` struct
- Improve tests coverage
    - Add test to check type resolution for elements with same _type-name_ and _package-name_ from another module
- Fix `GetElemKey` method not being able to get the type of the interface
- Fix an error where `interface` and `pointer` of the same type were being registered as the same type
- Create additional `"Do"` methods: **DoGet**, **DoGetGen**, **DoGetGenFunc**
- Refactor the **Storage/Injector** retrieval to return an `error` instead of a `bool`

## 20220724 - remy/v1.2.1

- Fix an error with `ReflectionOptions` not being applied to **StdInjector**
- Rename some internal attributes in storage

## 20220728 - remy/v1.3.0

- Improve test coverage
- Rename some internal attributes in storage
    - Test `SetGlobalInjector`
    - Test most generics utilities
- Return error on _internal register_ function
- Add `godoc` lines to **internal.types**
- Fix hidden error on `GetGen` function
- Fix error with bind register
    - Prevent overriding a same type with different type of bindings
- Cleanup `DependencyRetriever` methods
    - Removed unnecessary duplicate methods
- Internal improvements
    - Change `BindKey` type to prevent misplace errors
    - Remove duplicate use of **storage**

## 20220801 - remy/v1.4.0

- Create `CycleDetectorInjector` to be used in tests
    - Create a new error type
    - Create a new type in internal utilities
- Change use of unexported type to an exported in public pkg
    - remy public functions now use `Bind[T]` instead of `types.Bind[T]`
- Add `WrapRetriever` to **DependencyRetriever** interface
- Add panic recover to `Do` functions
- Remove `sync.RWMutex` from **globalInjector**
- Swap type `Injector` by `DependencyRetriever` in **Get** methods
- Boost performance by using pointer receiver in _Injector/Storage_ methods
