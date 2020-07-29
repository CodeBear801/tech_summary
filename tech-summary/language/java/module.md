
# java module

## Why
- just put needed package into final jar, solve the problem of small application's distribution, dll hell
- from jdk9 an java application must be packaged as java module.  JVM could check dependency at application's start time

Java module should
- declare what package is exported(visible)
- Circular Dependencies Not Allowed


## Example 

```bash
├── pl.tfij.java9modules.app
│   ├── module-info.java
│   └── pl
│       └── tfij
│           └── java9modules
│               └── app
│                   └── ModuleApp.java

```

compile
```bash
javac -d build --module-source-path src $(find src -name "*.java")
```

package
```bash
mkdir -p mods
jar --create --file=mods/pl.tfij.java9modules.app@1.0.jar --module-version=1.0 --main-class=pl.tfij.java9modules.app.ModuleApp -C build/pl.tfij.java9modules.app .
jar --create --file=mods/pl.tfij.java9modules.greetings@1.0.jar --module-version=1.0 -C build/pl.tfij.java9modules.greetings .
© 2020 GitHub, Inc.
```

link
```bash
jlink --module-path mods/:$JAVA_HOME/jmods --add-modules pl.tfij.java9modules.app --output pl.tfij.java9modules.app-image

# https://docs.oracle.com/javase/9/tools/jlink.htm#JSWOR-GUID-CECAC52B-CFEE-46CB-8166-F17A8E9280E9
# The jlink tool links a set of modules, along with their transitive dependences, to create a custom runtime image.

```

run
```bash
./pl.tfij.java9modules.app-image/bin/java -m pl.tfij.java9modules.app/pl.tfij.java9modules.app.ModuleApp

# ? what is java
```

## More info
- Java Modules http://tutorials.jenkov.com/java/modules.html#encapsulation-of-internal-packages
- Understanding Java 9 Modules https://www.oracle.com/corporate/features/understanding-java-9-modules.html
- http://chi.pl/2017/03/11/Quick-Introduction-to-Java9-Modularization.html
- https://github.com/tfij/Java-9-modules---the-simplest-example
