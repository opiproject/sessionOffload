/*
 *  Copyright (C) 2020 Palo Alto Networks Intellectual Property. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "opof_test.h"
#include "opof_stats_client.h"

using grpc::ClientReader;

void SessionStatsClient::getAllSessions(){
  sessionResponse response;
  statisticsRequestArgs request;
  ClientContext context;


  std::cout << "Getting all Sessions" << std::endl;

  std::unique_ptr<ClientReader <sessionResponse> > reader(
        stub_->getAllSessions(&context, request));
    while (reader->Read(&response)) {
      std::cout << "Found feature called "  << std::endl;
      std::cout << "Session ID: " << response.sessionid() << std::endl;
               // << feature.name() << " at "
               // << feature.location().latitude()/kCoordFactor_ << ", "
               // << feature.location().longitude()/kCoordFactor_ << std::endl;
    }

  Status status = reader->Finish();
  if (status.ok()) {
    std::cout << "getAllSessions rpc succeeded." << std::endl;
  } else {
    std::cout << "getAllSessions rpc failed." << std::endl;
  }
}
/*
Status status;
ClientContext context;
statisticsRequestArgs request;
actionType action;

request.set_inlif(s.inlif);
request.set_outlif(s.outlif);
request.set_sourceip(s.srcIP);
request.set_destinationip(s.dstIP);
request.set_ipversion((IP_VERSION)s.ipver);
request.set_sourceport(s.srcPort);
request.set_destinationport(s.dstPort);
request.set_protocolid((PROTOCOL_ID)s.proto);
action.set_actionvalue((ACTION_VALUE)s.actType);
action.set_actionnexthop(s.nextHop);
request.mutable_action()->CopyFrom(action);

std::cout << "Calling addSession " << std::endl;
sessionResponse statsResponse = stub_->getAllSessions(&context, request);
std::cout << "Called addSession " << std::endl;
if (status.ok()) {
    std::cout << "sessionID added is: " << response.sessionid() << std::endl;
    return "Success";
  } else {
    std::cout << status.error_code() << ": " << status.error_message()
              << std::endl;
    return "RPC failed";
  }
*/


// getSession
void SessionStatsClient::getClosedSessions(){
    sessionResponse response;
  statisticsRequestArgs request;
  ClientContext context;


  std::cout << "Getting Closed Sessions" << std::endl;

  std::unique_ptr<ClientReader <sessionResponse> > reader(
        stub_->getClosedSessions(&context, request));
    while (reader->Read(&response)) {
      std::cout << "getClosedSessions Called "  << std::endl;
      std::cout << "Session ID: " << response.sessionid() << std::endl;
               // << feature.name() << " at "
               // << feature.location().latitude()/kCoordFactor_ << ", "
               // << feature.location().longitude()/kCoordFactor_ << std::endl;
    }

  Status status = reader->Finish();
  if (status.ok()) {
    std::cout << "getClosedSessions rpc succeeded." << std::endl;
  } else {
    std::cout << "getClosedSessions rpc failed." << std::endl;
  }
}
/*
  sessionId sid;
  sessionResponse response;
  sid.set_sessionid(s.sessId);
  ClientContext context;
  std::cout << "Getting session ID: " << s.sessId << std::endl;
  sessionResponse statsResponse = stub_->getClosedSessions(&context, request);

  std::cout << "sessionID queried is: " << response.sessionid() << std::endl;
  std::cout << "session request status is: " << response.requeststatus() << std::endl;
  std::cout << "session state is: " << response.sessionstate() << std::endl;
  std::cout << "session close code is: " << response.sessionclosecode() << std::endl;
  std::cout << "session inPackets is: " << response.inpackets() << std::endl;
  std::cout << "session outPackets is: " << response.outpackets() << std::endl;

  if (status.ok()) {
    return "Success";
  } else {
    std::cout << status.error_code() << ": " << status.error_message()
              << std::endl;
    return "RPC failed";
  }
  

  //TODO: print timestamps
  //std::cout << "session starttime is: " << response.starttime() << std::endl;
  //std::cout << "session endtime is: " << response.endtime() << std::endl;
}

// deleteSession
std::string SessionTableClient::deleteSession(sessionTuple_t s){

sessionId sid;
sessionResponse response;
sid.set_sessionid(s.sessId);
ClientContext context;
Status status = stub_->deleteSession(&context, sid, &response);

std::cout << "sessionID deleted is: " << response.sessionid() << std::endl;
  std::cout << "session request status is: " << response.requeststatus() << std::endl;
  std::cout << "session state is: " << response.sessionstate() << std::endl;
  std::cout << "session close code is: " << response.sessionclosecode() << std::endl;
  std::cout << "session inPackets is: " << response.inpackets() << std::endl;
  std::cout << "session outPackets is: " << response.outpackets() << std::endl;

if (status.ok()) {
    return "Success";
  } else {
    std::cout << status.error_code() << ": " << status.error_message()
              << std::endl;
    return "RPC failed";
  }
  

  //TODO: print timestamps
  //std::cout << "session starttime is: " << response.starttime() << std::endl;
  //std::cout << "session endtime is: " << response.endtime() << std::endl;
  */

