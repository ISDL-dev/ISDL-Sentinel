import {
  Avatar,
  Box,
  Card,
  Checkbox,
  Divider,
  Flex,
  Table,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { Dispatch, SetStateAction, useMemo, useState } from "react";
import { GetUsersInfo200ResponseInner } from "../../schema";
import { GradeList, sortUsersByGrade } from "../../models/grade/grade";

interface UserListProps {
  userInfo: GetUsersInfo200ResponseInner[];
  targetUserId: number;
  setTargetUserId: Dispatch<SetStateAction<number>>;
}

export const UserList: React.FC<UserListProps> = ({
  userInfo,
  setTargetUserId,
  targetUserId,
}) => {
  const navigate = useNavigate();
  const changeUser = (userId: number) => {
    setTargetUserId(userId);
  };
  const [isShowObUser, setIsShowObUser] = useState(false);
  const sortedUserInfo = sortUsersByGrade(userInfo);
  const filteredUserInfo = useMemo(() => {
    return isShowObUser
      ? sortedUserInfo
      : sortedUserInfo.filter((user) => user.grade !== GradeList.OB);
  }, [sortedUserInfo, isShowObUser]);
  const handleCheckboxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const checked = e.target.checked;
    console.log("Checkbox checked:", checked);
    setIsShowObUser(checked);
  };
  return (
    <Card mb={{ base: "0px", lg: "20px" }} alignItems="center">
      <Box
        width="100%"
        height="100%"
        display="flex"
        flexDirection="column"
        justifyContent="space-between"
        overflowY="scroll"
      >
        <Box flex="1" m="15px">
          <Flex
            alignItems="center"
            justifyContent="space-between"
            width="100%"
            mb="10px"
          >
            <Text fontSize={{ base: "xl", md: "2xl" }} fontWeight="bold">
              ユーザ一覧
            </Text>
            <Flex alignItems="center" gap={3}>
              <Checkbox
                isChecked={isShowObUser}
                onChange={handleCheckboxChange}
              />
              <Text fontSize="sm" fontWeight="bold">
                卒業済みユーザの表示
              </Text>
            </Flex>
          </Flex>
          <Divider
            borderColor="rgba(79, 209, 197, 1)"
            borderWidth="3px"
            mb="10px"
          />
          <Box overflowY="auto" height={{ base: "25vh", md: "65vh" }}>
            <Table variant="simple" size={{ base: "sm", md: "md" }}>
              <Thead>
                <Tr>
                  <Th textAlign="center">ユーザ名</Th>
                  <Th textAlign="center">学年</Th>
                </Tr>
              </Thead>
              <Tbody>
                {filteredUserInfo.map((info) => (
                  <Tr
                    key={info.user_id}
                    onClick={() => changeUser(info.user_id)}
                    backgroundColor={
                      info.user_id === targetUserId ? "teal.100" : "transparent"
                    }
                    _hover={{
                      backgroundColor:
                        info.user_id === targetUserId ? "teal.200" : "gray.100",
                    }}
                  >
                    <Td textAlign="center">
                      <Flex alignItems={"center"} gap={3}>
                        <Box position="relative">
                          <Avatar
                            size={"md"}
                            src={info.avatar_img_path}
                            border="2px"
                            onClick={(e) => {
                              e.stopPropagation();
                              navigate("/profile", {
                                state: { userId: info.user_id },
                              });
                            }}
                          />
                        </Box>
                        {info.user_name}
                      </Flex>
                    </Td>
                    <Td textAlign="center">{info.grade}</Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </Box>
        </Box>
      </Box>
    </Card>
  );
};
