import { Box, Grid } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { UserList } from "../../features/UserList";
import { SettingInfo } from "../../features/SettingInfo";
import { settingApi } from "../../api";

// モック用
// backendの実装が完了したら消す
export type UserInfo = {
  user_id: number;
  name: string;
  mail_address: string;
  grade: string;
  avatar_img_path: string;
  role: string[];
};
const userInfo: UserInfo[] = [
  {
    user_id: 100,
    name: "酒部健太郎",
    mail_address: "sakabe.kentaro@mikilab.doshisha.ac.jp",
    grade: "M2",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["インフラ", "KC-111"],
  },
  {
    user_id: 300,
    name: "岡颯人",
    mail_address: "oka.hayato@mikilab.doshisha.ac.jp",
    grade: "M2",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["チーフ", "イベント"],
  },
  {
    user_id: 500,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 600,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 700,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 800,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 900,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 1000,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 1001,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
];
const roleListMock: string[] = [
  "チーフ",
  "メディア",
  "インフラ",
  "知的財産",
  "ミーティング",
  "Tex",
  "イベント",
  "KC-111",
];
const gradeListMock: string[] = [
  "Teacher",
  "D3",
  "D2",
  "D1",
  "M2",
  "M1",
  "U4",
  "OB",
];

export const UserSetting = () => {
  const [targetUserId, setTargetUserId] = useState(userInfo[0].user_id);
  const [roleList, setRoleList] = useState<string[]>([]);
  const [gradeList, setGradeList] = useState<string[]>([]);
  const fetchRoleList = async () => {
    // const roleResponse = await settingApi.getRoleName();
    setRoleList(roleListMock);
  };
  const fetchGradeList = async () => {
    // const roleResponse = await settingApi.getRoleName();
    setGradeList(gradeListMock);
  };
  useEffect(() => {
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
          base: "repeat(2, 1fr)",
          lg: "1fr",
        }}
        gap={{ base: "20px", xl: "20px" }}
      >
        <UserList
          userInfo={userInfo}
          setTargetUserId={setTargetUserId}
        ></UserList>
        <SettingInfo
          userInfo={userInfo}
          targetUserId={targetUserId}
          gradeList={gradeList}
          roleList={roleList}
        ></SettingInfo>
      </Grid>
    </Box>
  );
};
