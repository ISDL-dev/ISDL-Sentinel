import {
  Avatar,
  Box,
  Button,
  Card,
  Divider,
  Flex,
  Icon,
  IconButton,
  Select,
  Table,
  Tbody,
  Td,
  Text,
  Tr,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { RoleBadge } from "../RoleBadge";
import { EditIcon } from "@chakra-ui/icons";
import { ChangeEvent, useState } from "react";
import { FaCheck } from "react-icons/fa";
import { settingApi } from "../../api";
import { GetUsersInfo200ResponseInner } from "../../schema";

interface SettingInfoProps {
  userInfo: GetUsersInfo200ResponseInner[];
  targetUserId: number;
  roleList: string[];
  gradeList: string[];
  fetchUserList: () => Promise<void>;
}
const buttonWidth = {
  base: "100px",
  md: "150px",
};
export const SettingInfo: React.FC<SettingInfoProps> = ({
  userInfo,
  targetUserId,
  roleList,
  gradeList,
  fetchUserList,
}) => {
  const navigate = useNavigate();
  const targetUserInfo =
    userInfo.find((user) => user.user_id === targetUserId) ?? userInfo[0];
  const [changePendingUserInfo, setChangePendingUserInfo] =
    useState<GetUsersInfo200ResponseInner>(() => ({
      ...targetUserInfo,
      role_list: targetUserInfo.role_list || [],
    }));
  const [isEditing, setIsEditing] = useState(false);
  const handleGradeChange = (event: ChangeEvent<HTMLSelectElement>) => {
    const newInfo = { ...targetUserInfo };
    newInfo.grade = event.target.value;
    setChangePendingUserInfo(newInfo);
  };
  const handleRoleChange = (role: string) => {
    setChangePendingUserInfo((prevInfo) => {
      const newRoles = prevInfo.role_list.includes(role)
        ? prevInfo.role_list.filter((r) => r !== role)
        : [...prevInfo.role_list, role];
      return { ...prevInfo, role_list: newRoles };
    });
  };
  const checkRole = (role: string): boolean => {
    return changePendingUserInfo.role_list.includes(role);
  };
  const submitUserInfo = async () => {
    await settingApi.putUserById(changePendingUserInfo.user_id, {
      user_name: changePendingUserInfo.user_name,
      grade: changePendingUserInfo.grade,
      mail_address: changePendingUserInfo.mail_address,
      role_list: changePendingUserInfo.role_list,
    });
    fetchUserList();
  };
  return (
    <Card mb={{ base: "0px", lg: "20px" }} alignItems="center" p="30px">
      <Flex minHeight="15vh">
        <Flex
          alignItems="center"
          width="100%"
          maxWidth="800px"
          gap={12}
          mb="20px"
        >
          <Avatar
            size={"xl"}
            src={targetUserInfo?.avatar_img_path}
            border="2px"
            onClick={() =>
              navigate("/profile", {
                state: { userId: targetUserInfo?.user_id },
              })
            }
          />
          <Flex flexDirection="column" alignItems="flex-start" flex={1}>
            <Text fontWeight="bold" fontSize="2xl" mb="10px">
              {targetUserInfo?.user_name}
            </Text>
            <Text fontSize="md" fontWeight="400">
              メールアドレス: {targetUserInfo?.mail_address}
            </Text>
          </Flex>
        </Flex>
      </Flex>
      <Divider
        borderColor="rgba(79, 209, 197, 1)"
        borderWidth="3px"
        mb="30px"
      />
      <Flex flexDirection="column" gap={5}>
        <Flex alignItems="center" gap={10}>
          <Text fontWeight="bold" fontSize="xl" width="80px">
            学年:
          </Text>
          {isEditing ? (
            <Select
              value={changePendingUserInfo?.grade || undefined}
              onChange={(e) => {
                handleGradeChange(e);
              }}
              placeholder={targetUserInfo?.grade || "学年を選択"}
              css={{
                '& option[value=""]': {
                  display: "none",
                },
              }}
            >
              {gradeList.map((grade) => (
                <option key={grade} value={grade}>
                  {grade}
                </option>
              ))}
            </Select>
          ) : (
            <Text fontSize="xl" fontWeight="400">
              {targetUserInfo?.grade}
            </Text>
          )}
        </Flex>
        <Flex alignItems="center" gap={10}>
          <Text fontWeight="bold" fontSize="xl" width="80px">
            役割:
          </Text>
          <Flex flexWrap="wrap" gap={2}>
            {(
              (isEditing
                ? changePendingUserInfo?.role_list
                : targetUserInfo?.role_list) ?? []
            ).length > 0 ? (
              (isEditing
                ? changePendingUserInfo?.role_list
                : targetUserInfo?.role_list
              )?.map((role, index) => <RoleBadge key={index} text={role} />)
            ) : (
              <Text fontSize="xl" fontWeight="400">
                担当はありません
              </Text>
            )}
          </Flex>
        </Flex>
        {isEditing && (
          <Box
            borderRadius="md"
            borderWidth="2px"
            borderColor="gray.200"
            overflowY="auto"
            height="20vh"
            maxWidth="300px"
            mx="auto"
          >
            <Table variant="unstyled" size="sm">
              <Tbody>
                {roleList.map((role, index) => (
                  <Tr key={role} onClick={() => handleRoleChange(role)}>
                    <Td width="30px" px={2}>
                      {checkRole(role) && (
                        <Icon
                          as={FaCheck}
                          color="green.500"
                          boxSize={4}
                          ml={3}
                        />
                      )}
                    </Td>
                    <Td>
                      <RoleBadge key={index} text={role} />
                    </Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </Box>
        )}
      </Flex>
      {isEditing ? (
        <Flex position="absolute" bottom="2" right="2" gap={4} m={6}>
          <Button
            w={buttonWidth}
            colorScheme="gray"
            variant="solid"
            size="lg"
            onClick={() => {
              setChangePendingUserInfo({
                ...targetUserInfo,
                role_list: targetUserInfo.role_list || [],
              });
              setIsEditing(false);
            }}
          >
            キャンセル
          </Button>
          <Button
            w={buttonWidth}
            colorScheme="teal"
            variant="solid"
            size="lg"
            onClick={() => {
              setIsEditing(false);
              submitUserInfo();
            }}
          >
            更新
          </Button>
        </Flex>
      ) : (
        <IconButton
          aria-label="Change Avatar"
          icon={<EditIcon />}
          variant="ghost"
          colorScheme="teal"
          size="lg"
          position="absolute"
          bottom="2"
          right="2"
          fontSize="32px"
          m={6}
          onClick={() => {
            setChangePendingUserInfo({
              ...targetUserInfo,
              role_list: targetUserInfo.role_list || [],
            });
            setIsEditing(true);
          }}
        />
      )}
    </Card>
  );
};
