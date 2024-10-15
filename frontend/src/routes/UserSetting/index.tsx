import { Box, Grid } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { UserList } from "../../features/UserList";
import { SettingInfo } from "../../features/SettingInfo";
import { authenticationApi, settingApi } from "../../api";
import { GetUsersInfo200ResponseInner } from "../../schema";

export const UserSetting = () => {
  const [userInfoList, setUserInfoList] = useState<
    GetUsersInfo200ResponseInner[]
  >([]);
  const [targetUserId, setTargetUserId] = useState(0);
  const [roleList, setRoleList] = useState<string[]>([]);
  const [gradeList, setGradeList] = useState<string[]>([]);
  const fetchRoleList = async () => {
    const roleResponse = await settingApi.getRoleName();
    setRoleList(roleResponse.data);
  };
  const fetchGradeList = async () => {
    const gradeResponse = await authenticationApi.getGradeName();
    setGradeList(gradeResponse.data);
  };
  const fetchUserList = async () => {
    const userListResponse = await settingApi.getUsersInfo();
    setUserInfoList(userListResponse.data);
  };
  useEffect(() => {
    fetchUserList();
    fetchRoleList();
    fetchGradeList();
  }, []);
  return (
    <Box pt={{ base: "80px", md: "80px", xl: "10px" }}>
      <Grid
        templateColumns={{
          base: "1fr",
          lg: "0.6fr 1fr",
        }}
        templateRows={{
          base: "0.6fr 1.35fr",
          lg: "1fr",
        }}
        gap={{ base: "20px", xl: "20px" }}
      >
        <UserList
          userInfo={userInfoList}
          targetUserId={targetUserId}
          setTargetUserId={setTargetUserId}
        ></UserList>
        <SettingInfo
          userInfo={userInfoList}
          targetUserId={targetUserId}
          gradeList={gradeList}
          roleList={roleList}
          fetchUserList={fetchUserList}
        ></SettingInfo>
      </Grid>
    </Box>
  );
};
