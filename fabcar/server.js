'use strict';
const express = require('express');
const url = require('url');
const app = express();
const path = require('path');
const util = require('util');
const os = require('os');
const FabricClient = require('fabric-client');
const qs = require('qs');
const bodyParser = require('body-parser');
class Request {
  constructor(request, callback) {
    this.req = request;
    this.callback = callback;
    this.init();
    console.log(this.req, request);
  }
  init() {
    const fabricClient = new FabricClient();
    // setup the fabric network
    const channel = fabricClient.newChannel('mychannel');
    const peer = fabricClient.newPeer('grpc://localhost:7051');
    channel.addPeer(peer);
    var member_user = null;
    var store_path = path.join(__dirname, 'hfc-key-store');
    console.log('Store path: ' + store_path);
    var tx_id = null;

    // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    FabricClient.newDefaultKeyValueStore({
      path: store_path
    }).then((state_store) => {
      // assign the store to the fabric client
      fabricClient.setStateStore(state_store);
      var crypto_suite = FabricClient.newCryptoSuite();
      // use the same location for the state store (where the users' certificate are kept)
      // and the crypto store (where the users' keys are kept)
      var crypto_store = FabricClient.newCryptoKeyStore({
        path: store_path
      });
      crypto_suite.setCryptoKeyStore(crypto_store);
      fabricClient.setCryptoSuite(crypto_suite);

      // get the enrolled user from persistence, this user will sign all requests
      return fabricClient.getUserContext('user1', true);
    }).then((user_from_store) => {
      if (user_from_store && user_from_store.isEnrolled()) {
        console.log('Successfully loaded user1 from persistence');
        member_user = user_from_store;
      } else {
        this.callback({
          error: 'Failed to get user1.... run registerUser.js'
        });
        throw new Error('Failed to get user1.... run registerUser.js');
      }
      // queryCar chaincode function - requires 1 argument, ex: args: ['CAR4'],
      // queryAllCars chaincode function - requires no arguments , ex: args: [''],
      const request = {
        //targets : --- letting this default to the peers assigned to the channel
        chaincodeId: 'fabcar',
        fcn: this.req.fcn,
        args: this.req.args
      };
      // send the query proposal to the peer
      return channel.queryByChaincode(request);
    }).then((query_responses) => {
      console.log("Query has completed, checking results");
      // query_responses could have more than one  results if there multiple peers were used as targets
      if (query_responses && query_responses.length == 1) {
        if (query_responses[0] instanceof Error) {
          console.error("error from query = ", query_responses[0]);
          this.callback({
            error: query_responses[0]
          });
        } else {
          console.log("Response is ", query_responses[0].toString());
          this.callback({
            data: JSON.parse(query_responses[0].toString())
          });
        }
      } else {
        console.log("No payloads were returned from query");
        this.callback({
          data: ''
        });
      }
    }).catch((err) => {
      console.error('Failed to query successfully :: ' + err);
      this.callback({
        error: 'Failed to query successfully :: ' + err
      });
    });

  }
}

app.use(bodyParser.json());

app.post('/query', function (req, res) {
  let urlParts = req.body;
  const fcn = urlParts.fcn || '';
  const args = urlParts.args || [''];
  try {
    const request = new Request({
      fcn: fcn,
      args: args
    }, (response) => {
      res.json(response);
    });
  } catch (e) {
    console.error('catched exeption => ', e);
    res.send(e);
  }
});

app.listen(3000);
console.log('listen :3000');