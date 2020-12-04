package templates

const (
	MakeFile = `
hive:
	flutter packages pub run build_runner build
`

	MainCode = `import 'package:flutter/material.dart';

import 'package:hive/hive.dart';
import 'package:hive_flutter/hive_flutter.dart';

import 'screens/home.dart';
import 'screens/core/app.dart';
import 'screens/core/provider.dart';

void main() async {
  await Hive.initFlutter();

  runApp(App(child: Flutter()));
}

class Flutter extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var provider = Provider.of(context);
    return MaterialApp(
      home: HomeScreen(),
      theme: provider.config.theme
    );
  }
}
`


	HomeCode = `import 'package:flutter/material.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold();
  }
}
`

	AppCode = `import 'package:flutter/material.dart';

import 'provider.dart';

class App extends StatefulWidget {
  final Widget child;
  App({@required this.child});
  @override
  Provider createState() => Provider();
}

class Binding extends InheritedWidget {
  final Provider data;
  Binding({
    @required this.data,
    @required Widget child,
  }) : super(child: child);
  @override
  bool updateShouldNotify(Binding old) {
    return true;
  }
}
`

	ProviderCode = `import 'package:flutter/material.dart';

import 'app.dart';
import 'config.dart';
import 'storage.dart';

class Provider extends State<App> {
  Config config;
  Storage storage = Storage();

  set theme(ThemeData theme) => setState(() => config.theme = theme);

  @override
  void dispose() {
    super.dispose();
    storage.closeAllBox();
  }

  @override
  void initState() {
    storage.openAllBox();
    super.initState();
    config = Config(context);
  }

  @override
  Widget build(BuildContext context) =>
      Binding(data: this, child: widget.child);
  static Provider of(BuildContext context) =>
      context.dependOnInheritedWidgetOfExactType<Binding>().data;
}
`

ConfigCode = `
import 'package:flutter/material.dart';

final themeDark = ThemeData();
final themeDefault = ThemeData(
  appBarTheme: AppBarTheme(
    color: Colors.purple,
  ),
);

class Config {
  ThemeData theme = themeDefault;

  Config(BuildContext context) : size = MediaQuery.textScaleFactorOf(context);

  final double size;
  double get bar => size * 60;
  double get icon => size * 22;
  double get hint => size * 12;
}
`

StorageCode = `import 'dart:async';

import 'package:hive/hive.dart';

import 'data/example.dart';

class Storage {
  static void registerModels() {
    // Register all Adapters here
    Hive.registerAdapter(ExampleAdapter());
  }

  Storage();

  Box example;
  Completer<LazyBox> examples = Completer();

  Future<void> openAllBox() async {
    Storage.registerModels();

    example = await Hive.openBox<Example>("example");
    examples.complete(Hive.openLazyBox<Example>("examples"));
  }

  Future<void> closeAllBox() async {
    await Hive.close();
  }
}
`
ExampleCode = `
import 'package:hive/hive.dart';

part 'example.g.dart';

@HiveType(typeId: 1)
class Example {
  @HiveField(0)
  String example;
}
`

)
