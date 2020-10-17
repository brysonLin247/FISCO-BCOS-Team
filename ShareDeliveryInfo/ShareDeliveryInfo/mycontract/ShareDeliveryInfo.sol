pragma solidity >= 0.4.25;

contract ShareDeliveryInfo {
    uint package_uid; //包裹的UID

    mapping(uint64 => string) public delivery_path;//记录快递到达站点路径

    constructor() public {
    }

    //获取快递传送路径信息
    function get(uint64 _pkg_uid) public view returns(string) {
        return delivery_path[_pkg_uid];
    }

    //快递到达新站点，更新信息
    function set(uint64 _pkg_uid, string station) public {
        delivery_path[_pkg_uid] = strConcat(delivery_path[_pkg_uid], "； 到达站点：");
        delivery_path[_pkg_uid] = strConcat(delivery_path[_pkg_uid], station);
    }

    //用于拼接字符串
    function strConcat(string _a, string _b) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory ret = new string(_ba.length + _bb.length);
        bytes memory bret = bytes(ret);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++)bret[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) bret[k++] = _bb[i];
        return string(ret);
   }
}