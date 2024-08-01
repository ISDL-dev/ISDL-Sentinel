import { Box, Table, Thead, Tbody, Tr, Th, Td, Image, Text, Divider, IconButton } from '@chakra-ui/react';
import { DeleteIcon, EditIcon } from '@chakra-ui/icons';
import Card from '../../features/Card';
import { GetUserById200ResponseAvatarListInner } from '../../schema';

interface AvatarListProps {
  avatars: GetUserById200ResponseAvatarListInner[];
  onAvatarClick: (avatarId: number) => void;
  onDeleteClick: (avatarId: number) => void;
  [x: string]: any;
}

const AvatarList: React.FC<AvatarListProps> = ({ avatars, onAvatarClick, onDeleteClick, ...rest }) => {
  return (
    <Card mb={{ base: '0px', lg: '20px' }} alignItems='center' {...rest}>
      <Box width="100%">
        <Text fontSize="2xl" fontWeight="bold" mb="20px" textAlign="left">
          Avatar List
        </Text>
        <Divider borderColor="rgba(79, 209, 197, 1)" borderWidth="3px" mb="20px" />
        <Box overflowX="auto">
          <Table variant="simple" size={{ base: 'sm', md: 'md' }}>
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
                    <Image 
                      src={`./avatar/${avatar.img_path}`} 
                      boxSize={{ base: '40px', md: '50px' }}
                    />
                  </Td>
                  <Td textAlign="center">
                    <Box display="flex" justifyContent="center">
                      <IconButton 
                        aria-label="Change Avatar"
                        icon={<EditIcon />}
                        onClick={() => onAvatarClick(avatar.avatar_id)}
                        variant="ghost"
                        colorScheme="teal"
                        size={{ base: 'sm', md: 'md' }}
                        mr={{ base: '1', md: '2' }}
                      />
                      {avatar.avatar_id !== 1 && (
                        <IconButton 
                          aria-label="Delete Avatar"
                          icon={<DeleteIcon />}
                          onClick={() => onDeleteClick(avatar.avatar_id)}
                          variant="ghost"
                          colorScheme="teal"
                          size={{ base: 'sm', md: 'md' }}
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
    </Card>
  );
};

export default AvatarList;
