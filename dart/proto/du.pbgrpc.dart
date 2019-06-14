///
//  Generated code. Do not modify.
//  source: du.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core show int, String, List;

import 'package:grpc/service_api.dart' as $grpc;
import 'du.pb.dart' as $0;
export 'du.pb.dart';

class duClient extends $grpc.Client {
  static final _$hi = $grpc.ClientMethod<$0.Name, $0.Reply>(
      '/du/hi',
      ($0.Name value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Reply.fromBuffer(value));

  duClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.Reply> hi($0.Name request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$hi, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class duServiceBase extends $grpc.Service {
  $core.String get $name => 'du';

  duServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Name, $0.Reply>(
        'hi',
        hi_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Name.fromBuffer(value),
        ($0.Reply value) => value.writeToBuffer()));
  }

  $async.Future<$0.Reply> hi_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return hi(call, await request);
  }

  $async.Future<$0.Reply> hi($grpc.ServiceCall call, $0.Name request);
}
