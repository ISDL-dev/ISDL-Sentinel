import {
  Box,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Text,
  Divider,
  IconButton,
  Button,
  Input,
  Avatar,
} from "@chakra-ui/react";
import { DeleteIcon, EditIcon } from "@chakra-ui/icons";
import Card from "../../features/Card";
import { GetUserById200ResponseAvatarListInner } from "../../schema";

interface AvatarListProps {
  avatars: GetUserById200ResponseAvatarListInner[];
  onAvatarClick: (avatarId: number) => void;
  onDeleteClick: (avatarId: number) => void;
  selectedFile: File | null;
  setSelectedFile: (file: File | null) => void;
  onSaveUpload: () => void;
  onCancelUpload: () => void;
  [x: string]: any;
}

const AvatarList: React.FC<AvatarListProps> = ({
  avatars,
  onAvatarClick,
  onDeleteClick,
  selectedFile,
  setSelectedFile,
  onSaveUpload,
  onCancelUpload,
  ...rest
}) => {
  return (
    <Card mb={{ base: "0px", lg: "20px" }} alignItems="center" {...rest}>
      <Box
        width="100%"
        height="100%"
        display="flex"
        flexDirection="column"
        justifyContent="space-between"
      >
        <Box overflowY="auto" flex="1" mb="10px">
          <Text fontSize="2xl" fontWeight="bold" mb="20px" textAlign="left">
            Avatar List
          </Text>
          <Divider
            borderColor="rgba(79, 209, 197, 1)"
            borderWidth="3px"
            mb="20px"
          />

          <Box overflowX="auto" mb="20px">
            <Table variant="simple" size={{ base: "sm", md: "md" }}>
              <Thead>
                <Tr>
                  <Th textAlign="center">Index</Th>
                  <Th textAlign="center">Avatar Image</Th>
                  <Th textAlign="center">Actions</Th>
                </Tr>
              </Thead>
              <Tbody>
                {avatars.map((avatar, index) => (
                  <Tr key={avatar.avatar_id}>
                    <Td textAlign="center">{index + 1}</Td>
                    <Td textAlign="center">
                      <Box
                        display="flex"
                        justifyContent="center"
                        alignItems="center"
                        height={{ base: "40px", md: "50px" }}
                      >
                        <Avatar size="md" src={avatar.img_path} border="2px" />
                      </Box>
                    </Td>
                    <Td textAlign="center">
                      <Box display="flex" justifyContent="center">
                        <IconButton
                          aria-label="Change Avatar"
                          icon={<EditIcon />}
                          onClick={() => onAvatarClick(avatar.avatar_id)}
                          variant="ghost"
                          colorScheme="teal"
                          size={{ base: "sm", md: "md" }}
                          mr={{ base: "1", md: "2" }}
                        />
                        {avatar.avatar_id !== 1 && (
                          <IconButton
                            aria-label="Delete Avatar"
                            icon={<DeleteIcon />}
                            onClick={() => onDeleteClick(avatar.avatar_id)}
                            variant="ghost"
                            colorScheme="teal"
                            size={{ base: "sm", md: "md" }}
                          />
                        )}
                      </Box>
                    </Td>
                  </Tr>
                ))}
              </Tbody>
            </Table>
          </Box>
        </Box>

        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          flexDirection={{ base: "column", sm: "row" }}
          position="sticky"
          bottom="0"
          left="0"
          width="100%"
          p="10px"
          bg="white"
          borderColor="gray.200"
          gap={2}
        >
          {selectedFile ? (
            <>
              <Box display="flex" alignItems="center" gap="10" mt="4">
                <Avatar
                  size="md"
                  src={URL.createObjectURL(selectedFile)}
                  border="2px"
                />
                <Button
                  colorScheme="teal"
                  onClick={onSaveUpload}
                  width={{ base: "100%", sm: "auto" }}
                >
                  Save
                </Button>
                <Button
                  variant="outline"
                  onClick={onCancelUpload}
                  width={{ base: "100%", sm: "auto" }}
                >
                  Cancel
                </Button>
              </Box>
            </>
          ) : (
            <>
              <Input
                type="file"
                accept="image/*"
                onChange={(e) => {
                  if (e.target.files && e.target.files[0]) {
                    setSelectedFile(e.target.files[0]);
                  }
                }}
                display="none"
                id="avatar-upload"
              />
              <Button
                as="label"
                htmlFor="avatar-upload"
                colorScheme="teal"
                width={{ base: "100%", sm: "auto" }}
              >
                Upload
              </Button>
            </>
          )}
        </Box>
      </Box>
    </Card>
  );
};

export default AvatarList;
