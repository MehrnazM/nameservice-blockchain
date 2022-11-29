import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteName } from "./types/nameservice/nameservice/tx";
import { MsgBuyName } from "./types/nameservice/nameservice/tx";
import { MsgSetName } from "./types/nameservice/nameservice/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/nameservice.nameservice.MsgDeleteName", MsgDeleteName],
    ["/nameservice.nameservice.MsgBuyName", MsgBuyName],
    ["/nameservice.nameservice.MsgSetName", MsgSetName],
    
];

export { msgTypes }