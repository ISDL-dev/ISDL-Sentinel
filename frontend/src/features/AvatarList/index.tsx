import { Box, Table, Thead, Tbody, Tr, Th, Td, Image, Text, Divider } from '@chakra-ui/react';
import Card from '../../features/Card';
import { GetUserById200ResponseAvatarListInner } from '../../schema';

interface AvatarListProps {
  avatars: GetUserById200ResponseAvatarListInner[];
  onAvatarClick: (avatarId: number) => void;
  [x: string]: any;
}

const AvatarList: React.FC<AvatarListProps> = ({ avatars, onAvatarClick, ...rest }) => {
  return (
    <Card mb={{ base: '0px', lg: '20px' }} alignItems='center' {...rest}>
      <Box width="100%">
        <Text fontSize="2xl" fontWeight="bold" mb="20px" textAlign="left">
          Avatar List
        </Text>
        <Divider borderColor="rgba(79, 209, 197, 1)" borderWidth="3px" mb="20px" />
        <Box overflowX="auto">
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th textAlign="center">Avatar Image</Th>
                <Th textAlign="center">Avatar Name</Th>
                <Th textAlign="center">Rarity</Th>
              </Tr>
            </Thead>
            <Tbody>
              {avatars.map((avatar) => (
                <Tr key={avatar.avatar_id}>
                  <Td textAlign="center">
                    <Image 
                      src={`./avatar/${avatar.img_path}`} 
                      alt={avatar.avatar_name} 
                      boxSize="50px" 
                      cursor="pointer"
                      onClick={() => onAvatarClick(avatar.avatar_id)}
                    />
                  </Td>
                  <Td textAlign="center">{avatar.avatar_name}</Td>
                  <Td textAlign="center">{avatar.rarity}</Td>
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
