import { Box, Grid } from "@chakra-ui/react";
import { useState } from "react";
import { UserList } from "../../features/UserList";
import { SettingInfo } from "../../features/SettingInfo";

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
    user_id: 1,
    name: "酒部健太郎",
    mail_address: "sakabe.kentaro@mikilab.doshisha.ac.jp",
    grade: "M2",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra", "kc111"],
  },
  {
    user_id: 3,
    name: "岡颯人",
    mail_address: "oka.hayato@mikilab.doshisha.ac.jp",
    grade: "M2",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["chief", "event"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
  {
    user_id: 5,
    name: "花本凪",
    mail_address: "hanamoto.nagi@mikilab.doshisha.ac.jp",
    grade: "OB",
    avatar_img_path:
      "https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000",
    role: ["infra"],
  },
];

export const UserSetting = () => {
  const [targetUserId, setTargetUserId] = useState(userInfo[0].user_id);
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
        ></SettingInfo>
      </Grid>
    </Box>
  );
};
