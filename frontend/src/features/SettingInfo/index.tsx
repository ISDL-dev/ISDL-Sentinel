import {
  Avatar,
  Card,
  Divider,
  Flex,
  IconButton,
  Text,
} from "@chakra-ui/react";
import { UserInfo } from "../../routes/UserSetting";
import { useNavigate } from "react-router-dom";
import { RoleBadge } from "../RoleBadge";
import { EditIcon } from "@chakra-ui/icons";

interface SettingInfoProps {
  userInfo: UserInfo[];
  targetUserId: number;
}
export const SettingInfo: React.FC<SettingInfoProps> = ({
  userInfo,
  targetUserId,
}) => {
  const navigate = useNavigate();
  const targetUserInfo = userInfo.find((user) => user.user_id === targetUserId);
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
              {targetUserInfo?.name}
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
          <Text fontSize="xl" fontWeight="400">
            {targetUserInfo?.grade}
          </Text>
        </Flex>
        <Flex alignItems="center" gap={10}>
          <Text fontWeight="bold" fontSize="xl" width="80px">
            役割:
          </Text>
          <Flex flexWrap="wrap" gap={2}>
            {targetUserInfo && targetUserInfo.role.length > 0 ? (
              targetUserInfo?.role.map((role, index) => (
                <RoleBadge key={index} text={role} />
              ))
            ) : (
              <Text fontSize="xl" fontWeight="400">
                担当はありません
              </Text>
            )}
          </Flex>
        </Flex>
      </Flex>
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
      />
    </Card>
  );
};
