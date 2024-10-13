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
  useToast,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { RoleBadge } from "../RoleBadge";
import { EditIcon } from "@chakra-ui/icons";
import { ChangeEvent, useState } from "react";
import { FaCheck } from "react-icons/fa";
import { settingApi } from "../../api";
import { GetUsersInfo200ResponseInner } from "../../schema";
import { useUser } from "../../userContext";
import { isChief } from "../../models/role/role";

interface SettingInfoProps {
  userInfo: GetUsersInfo200ResponseInner[];
  targetUserId: number;
  roleList: string[];
  gradeList: string[];
  fetchUserList: () => Promise<void>;
}
const buttonWidth = {
  base: "80px",
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
  const toast = useToast();
  const { authUser } = useUser();
  const targetUserInfo =
    userInfo.find((user) => user.user_id === targetUserId) ?? userInfo[0];
  const [changePendingUserInfo, setChangePendingUserInfo] =
    useState<GetUsersInfo200ResponseInner>(() => ({
      ...targetUserInfo,
      role_list: targetUserInfo?.role_list ?? [],
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
    try {
      await settingApi.putUserById(changePendingUserInfo.user_id, {
        user_name: changePendingUserInfo.user_name,
        grade: changePendingUserInfo.grade,
        mail_address: changePendingUserInfo.mail_address,
        role_list: changePendingUserInfo.role_list,
      });
      fetchUserList();
      toast({
        title: "ユーザ情報が正常に更新されました。",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
    } catch (err) {
      toast({
        title: "ユーザ情報の更新に失敗しました。",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
    }
  };
  return (
    <Card
      mb={{ base: "0px", lg: "20px" }}
      alignItems="center"
      p={{ base: "15px", md: "30px" }}
    >
      <Flex minHeight={{ base: "10vh", md: "15vh" }}>
        <Flex
          alignItems="center"
          justifyContent="center"
          width="100%"
          maxWidth={{ base: "100%", md: "800px" }}
          gap={12}
          mb="20px"
          ml={{ base: "20px", md: "0px" }}
          flexDirection="row"
        >
          <Avatar
            size={{ base: "lg", md: "xl" }}
            src={targetUserInfo?.avatar_img_path}
            border="2px"
            onClick={() =>
              navigate("/profile", {
                state: { userId: targetUserInfo?.user_id },
              })
            }
          />
          <Flex flexDirection="column" alignItems="flex-start" flex={1}>
            <Text
              fontWeight="bold"
              fontSize={{ base: "xl", md: "2xl" }}
              mb="10px"
            >
              {targetUserInfo?.user_name}
            </Text>
            <Text fontSize={{ base: "sm", md: "md" }} fontWeight="400">
              メールアドレス: {targetUserInfo?.mail_address}
            </Text>
          </Flex>
        </Flex>
      </Flex>
      <Divider
        borderColor="rgba(79, 209, 197, 1)"
        borderWidth="3px"
        mb={{ base: "10px", md: "30px" }}
      />
      <Flex flexDirection="column" gap={{ base: 2, md: 5 }}>
        <Flex alignItems="center" gap={{ base: 5, md: 10 }}>
          <Text
            fontWeight="bold"
            fontSize={{ base: "sm", md: "xl" }}
            width="80px"
          >
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
            <Text fontSize={{ base: "sm", md: "xl" }} fontWeight="400">
              {targetUserInfo?.grade}
            </Text>
          )}
        </Flex>
        <Flex alignItems="center" gap={{ base: 5, md: 10 }}>
          <Text
            fontWeight="bold"
            fontSize={{ base: "sm", md: "xl" }}
            width="80px"
          >
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
              <Text fontSize={{ base: "sm", md: "xl" }} fontWeight="400">
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
            height={{ base: "15vh", md: "20vh" }}
            maxWidth="300px"
            mx="auto"
          >
            <Table variant="unstyled" size={{ base: "xs", md: "sm" }}>
              <Tbody>
                {roleList.map((role, index) => (
                  <Tr key={role} onClick={() => handleRoleChange(role)}>
                    <Td width="30px" height={{ base: "10px" }} px={2}>
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
      {isChief(authUser) &&
        (isEditing ? (
          <Flex
            position="absolute"
            bottom="2"
            right="2"
            gap={4}
            m={{ base: 3, md: 6 }}
          >
            <Button
              w={buttonWidth}
              colorScheme="gray"
              variant="solid"
              size={{ base: "xs", md: "lg" }}
              fontSize={{ base: "xs", md: "md" }}
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
              size={{ base: "xs", md: "lg" }}
              fontSize={{ base: "xs", md: "md" }}
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
        ))}
    </Card>
  );
};
