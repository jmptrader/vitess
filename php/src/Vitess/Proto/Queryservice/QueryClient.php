<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: queryservice.proto

namespace Vitess\Proto\Queryservice {

  class QueryClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts, $channel = null) {
      parent::__construct($hostname, $opts, $channel);
    }
    /**
     * @param Vitess\Proto\Query\ExecuteRequest $input
     */
    public function Execute(\Vitess\Proto\Query\ExecuteRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/Execute', $argument, '\Vitess\Proto\Query\ExecuteResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\ExecuteBatchRequest $input
     */
    public function ExecuteBatch(\Vitess\Proto\Query\ExecuteBatchRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/ExecuteBatch', $argument, '\Vitess\Proto\Query\ExecuteBatchResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\StreamExecuteRequest $input
     */
    public function StreamExecute($argument, $metadata = array(), $options = array()) {
      return $this->_serverStreamRequest('/queryservice.Query/StreamExecute', $argument, '\Vitess\Proto\Query\StreamExecuteResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\BeginRequest $input
     */
    public function Begin(\Vitess\Proto\Query\BeginRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/Begin', $argument, '\Vitess\Proto\Query\BeginResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\CommitRequest $input
     */
    public function Commit(\Vitess\Proto\Query\CommitRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/Commit', $argument, '\Vitess\Proto\Query\CommitResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\RollbackRequest $input
     */
    public function Rollback(\Vitess\Proto\Query\RollbackRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/Rollback', $argument, '\Vitess\Proto\Query\RollbackResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\BeginExecuteRequest $input
     */
    public function BeginExecute(\Vitess\Proto\Query\BeginExecuteRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/BeginExecute', $argument, '\Vitess\Proto\Query\BeginExecuteResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\BeginExecuteBatchRequest $input
     */
    public function BeginExecuteBatch(\Vitess\Proto\Query\BeginExecuteBatchRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/BeginExecuteBatch', $argument, '\Vitess\Proto\Query\BeginExecuteBatchResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\SplitQueryRequest $input
     */
    public function SplitQuery(\Vitess\Proto\Query\SplitQueryRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/queryservice.Query/SplitQuery', $argument, '\Vitess\Proto\Query\SplitQueryResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Query\StreamHealthRequest $input
     */
    public function StreamHealth($argument, $metadata = array(), $options = array()) {
      return $this->_serverStreamRequest('/queryservice.Query/StreamHealth', $argument, '\Vitess\Proto\Query\StreamHealthResponse::deserialize', $metadata, $options);
    }
  }
}