#!/usr/bin/env dart
import 'package:grpc/grpc.dart';

import 'proto/du.pb.dart';
import 'proto/du.pbgrpc.dart';

class Client {
  ClientChannel channel;
  duClient stub;

  Future<void> main(List<String> args) async {
    channel = ClientChannel('127.0.0.1',
        port: 50051,
        options:
            const ChannelOptions(credentials: ChannelCredentials.insecure()));
    stub =
        duClient(channel, options: CallOptions(timeout: Duration(seconds: 30)));
    // Run all of the demos in order.
    try {
      await hi();
    } catch (e) {
      print('Caught error: $e');
    }
    await channel.shutdown();
  }

  /// Run the getFeature demo. Calls getFeature with a point known to have a
  /// feature and a point known not to have a feature.
  Future<void> hi() async {
    for (var i in Iterable.generate(10)) {
      final name = Name()..name = "鹏 $i";
      Reply reply = await stub.hi(name);
      print(reply.message);
    }
  }
}

main(List<String> args) async {
  await Client().main(args);
}
