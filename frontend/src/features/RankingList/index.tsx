import {
  Avatar,
  AvatarBadge,
  Button,
  Flex,
  Grid,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";

export const RankingList = (placeholder: { placeholder: string }) => {
  return (
    <TableContainer
      pb={14}
      pr={14}
      pl={14}
      outlineOffset={2}
      overflowX="unset"
      overflowY="scroll"
      height="65vh"
    >
      <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
        <Thead top={0}>
          <Tr bgColor="#E6EBED">
            <Th w="33%">順位</Th>
            <Th w="33%">ユーザ</Th>
            <Th w="33%">{placeholder.placeholder}</Th>
          </Tr>
        </Thead>
        <Tbody outline="1px"></Tbody>
      </Table>
    </TableContainer>
  );
};
