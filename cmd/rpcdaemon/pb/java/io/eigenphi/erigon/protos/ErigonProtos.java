// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

public final class ErigonProtos {
  private ErigonProtos() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_stackFrame_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_stackFrame_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_InternalTxn_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_InternalTxn_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_EventLog_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_EventLog_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_Receipt_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_Receipt_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_Transaction_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_Transaction_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_Block_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_Block_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_erigon_TraceTransaction_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_erigon_TraceTransaction_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\013Block.proto\022\006erigon\"\254\001\n\nstackFrame\022\014\n\004" +
      "type\030\001 \001(\t\022\r\n\005label\030\002 \001(\t\022\014\n\004from\030\003 \001(\t\022" +
      "\n\n\002to\030\004 \001(\t\022\027\n\017contractCreated\030\005 \001(\t\022\r\n\005" +
      "value\030\006 \001(\t\022\r\n\005input\030\007 \001(\t\022\r\n\005error\030\010 \001(" +
      "\t\022!\n\005calls\030\t \003(\0132\022.erigon.stackFrame\"O\n\013" +
      "InternalTxn\022\027\n\017transactionHash\030\001 \001(\t\022\014\n\004" +
      "from\030\002 \001(\t\022\n\n\002to\030\003 \001(\t\022\r\n\005value\030\004 \001(\t\"\242\002" +
      "\n\010EventLog\022\r\n\005chain\030\001 \001(\t\022\017\n\007address\030\002 \001" +
      "(\t\022\016\n\006topic0\030\003 \001(\t\022\016\n\006topic1\030\004 \001(\t\022\016\n\006to" +
      "pic2\030\005 \001(\t\022\016\n\006topic3\030\006 \001(\t\022\014\n\004data\030\007 \001(\t" +
      "\022\023\n\013blockNumber\030\010 \001(\003\022\026\n\016blockTimestamp\030" +
      "\t \001(\003\022\027\n\017transactionHash\030\n \001(\t\022\030\n\020transa" +
      "ctionIndex\030\013 \001(\005\022\021\n\tblockHash\030\014 \001(\t\022\020\n\010l" +
      "ogIndex\030\r \001(\005\022\017\n\007removed\030\016 \001(\010\022\022\n\nsender" +
      "Info\030\017 \001(\t\"\244\002\n\007Receipt\022\r\n\005chain\030\001 \001(\t\022\027\n" +
      "\017transactionHash\030\002 \001(\t\022\030\n\020transactionInd" +
      "ex\030\003 \001(\005\022\021\n\tblockHash\030\004 \001(\t\022\023\n\013blockNumb" +
      "er\030\005 \001(\003\022\017\n\007gasUsed\030\006 \001(\003\022\027\n\017contractAdd" +
      "ress\030\007 \001(\t\022 \n\030transactionReceiptStatus\030\010" +
      " \001(\010\022\022\n\neventCount\030\t \001(\005\022\026\n\016blockTimesta" +
      "mp\030\n \001(\003\022#\n\teventLogs\030\013 \003(\0132\020.erigon.Eve" +
      "ntLog\022\022\n\nsenderInfo\030\014 \001(\t\"\276\001\n\013Transactio" +
      "n\022\027\n\017transactionHash\030\001 \001(\t\022\030\n\020transactio" +
      "nIndex\030\002 \001(\005\022\023\n\013blockNumber\030\003 \001(\003\022\023\n\013fro" +
      "mAddress\030\004 \001(\t\022\021\n\ttoAddress\030\005 \001(\t\022\r\n\005non" +
      "ce\030\006 \001(\003\022\030\n\020transactionValue\030\007 \001(\t\022\026\n\016bl" +
      "ockTimestamp\030\010 \001(\003\"\240\001\n\005Block\022\023\n\013blockNum" +
      "ber\030\002 \001(\003\022\021\n\tblockHash\030\003 \001(\t\022\022\n\nparentHa" +
      "sh\030\004 \001(\t\022\r\n\005miner\030\006 \001(\t\022\021\n\tblockSize\030\007 \001" +
      "(\005\022\020\n\010gasLimit\030\010 \001(\003\022\017\n\007gasUsed\030\t \001(\003\022\026\n" +
      "\016blockTimestamp\030\n \001(\003\"\357\001\n\020TraceTransacti" +
      "on\022\023\n\013blockNumber\030\001 \001(\003\022\027\n\017transactionHa" +
      "sh\030\002 \001(\t\022\030\n\020transactionIndex\030\003 \001(\005\022\023\n\013fr" +
      "omAddress\030\004 \001(\t\022\021\n\ttoAddress\030\005 \001(\t\022\020\n\010ga" +
      "sPrice\030\006 \001(\003\022\r\n\005input\030\007 \001(\t\022\r\n\005nonce\030\010 \001" +
      "(\003\022\030\n\020transactionValue\030\t \001(\t\022!\n\005stack\030\n " +
      "\001(\0132\022.erigon.stackFrameB7\n\031io.eigenphi.e" +
      "rigon.protosB\014ErigonProtosP\001Z\n./protobuf" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_erigon_stackFrame_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_erigon_stackFrame_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_stackFrame_descriptor,
        new java.lang.String[] { "Type", "Label", "From", "To", "ContractCreated", "Value", "Input", "Error", "Calls", });
    internal_static_erigon_InternalTxn_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_erigon_InternalTxn_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_InternalTxn_descriptor,
        new java.lang.String[] { "TransactionHash", "From", "To", "Value", });
    internal_static_erigon_EventLog_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_erigon_EventLog_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_EventLog_descriptor,
        new java.lang.String[] { "Chain", "Address", "Topic0", "Topic1", "Topic2", "Topic3", "Data", "BlockNumber", "BlockTimestamp", "TransactionHash", "TransactionIndex", "BlockHash", "LogIndex", "Removed", "SenderInfo", });
    internal_static_erigon_Receipt_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_erigon_Receipt_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_Receipt_descriptor,
        new java.lang.String[] { "Chain", "TransactionHash", "TransactionIndex", "BlockHash", "BlockNumber", "GasUsed", "ContractAddress", "TransactionReceiptStatus", "EventCount", "BlockTimestamp", "EventLogs", "SenderInfo", });
    internal_static_erigon_Transaction_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_erigon_Transaction_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_Transaction_descriptor,
        new java.lang.String[] { "TransactionHash", "TransactionIndex", "BlockNumber", "FromAddress", "ToAddress", "Nonce", "TransactionValue", "BlockTimestamp", });
    internal_static_erigon_Block_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_erigon_Block_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_Block_descriptor,
        new java.lang.String[] { "BlockNumber", "BlockHash", "ParentHash", "Miner", "BlockSize", "GasLimit", "GasUsed", "BlockTimestamp", });
    internal_static_erigon_TraceTransaction_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_erigon_TraceTransaction_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_erigon_TraceTransaction_descriptor,
        new java.lang.String[] { "BlockNumber", "TransactionHash", "TransactionIndex", "FromAddress", "ToAddress", "GasPrice", "Input", "Nonce", "TransactionValue", "Stack", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}