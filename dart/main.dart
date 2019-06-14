import 'dart:async';
import 'dart:math' show Random;

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

  void printReply(Reply reply) {
    print(reply.message);
  }

  /// Run the getFeature demo. Calls getFeature with a point known to have a
  /// feature and a point known not to have a feature.
  Future<void> hi() async {
    final name = Name()..name = "鹏";

    printReply(await stub.hi(name));
    printReply(await stub.hi(name));
    printReply(await stub.hi(name));
    printReply(await stub.hi(name));
  }
}

main(List<String> args) async {
  await Client().main(args);
}
