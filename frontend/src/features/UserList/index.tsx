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
import { UserInfo } from "../../routes/UserSetting";
import { useNavigate } from "react-router-dom";
import { Dispatch, SetStateAction } from "react";

interface UserListProps {
  userInfo: UserInfo[];
  setTargetUserId: Dispatch<SetStateAction<number>>;
}

export const UserList: React.FC<UserListProps> = ({
  userInfo,
  setTargetUserId,
}) => {
  const navigate = useNavigate();
  const changeUser = (userId: number) => {
    setTargetUserId(userId);
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
            <Text fontSize="2xl" fontWeight="bold">
              ユーザ一覧
            </Text>
            <Flex alignItems="center" gap={3}>
              <Checkbox type="checkbox" />
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
          <Box overflowY="auto" maxHeight="65vh">
            <Table variant="simple" size={{ base: "sm", md: "md" }}>
              <Thead>
                <Tr>
                  <Th textAlign="center">ユーザ名</Th>
                  <Th textAlign="center">学年</Th>
                </Tr>
              </Thead>
              <Tbody>
                {userInfo.map((info, index) => (
                  <Tr
                    key={info.user_id}
                    onClick={() => changeUser(info.user_id)}
                  >
                    <Td textAlign="center">
                      <Flex alignItems={"center"} gap={3}>
                        <Box position="relative">
                          <Avatar
                            size={"md"}
                            src={info.avatar_img_path}
                            border="2px"
                            onClick={() =>
                              navigate("/profile", {
                                state: { userId: info.user_id },
                              })
                            }
                          />
                        </Box>
                        {info.name}
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
