import {
  Avatar,
  AvatarBadge,
  Button,
  Flex,
  Grid,
  Table,
  TableCellProps,
  TableColumnHeaderProps,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import "./Home.css";
import dayjs from "dayjs";
import "dayjs/locale/ja";

dayjs.locale("ja");
const attendees = [
  {
    id: 1,
    name: "酒部健太郎",
    placeName: "KC104",
    enteredAt: new Date(2024, 6, 13),
    avaterPath:
      "https://images.unsplash.com/photo-1619946794135-5bc917a27793?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9",
  },
  {
    id: 3,
    name: "岡颯人",
    placeName: "KC104",
    enteredAt: new Date(2024, 6, 14),
    avaterPath:
      "https://images.unsplash.com/photo-1619946794135-5bc917a27793?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9",
  },
];

const decodeDate = (date: Date) => {
  return `${dayjs(date).format("MM月DD日")}（${dayjs(date).format(
    "ddd"
  )}）${dayjs(date).format("HH時MM分")}`;
};

function Home() {
  return (
    <div>
      <Grid
        templateColumns="repeat(3, 1fr)"
        alignItems={"center"}
        w={"-moz-max-content"}
      >
        <h1 className="block mb-1 text-4xl font-bold text-gray-900 dark:text-white p-3 text-left">
          出席者一覧
        </h1>
        <Grid
          templateColumns="repeat(2, 1fr)"
          alignItems={"center"}
          w={"-moz-max-content"}
          column={3}
        >
          <Button colorScheme="teal" variant="solid" size="lg" width={36}>
            入室
          </Button>
          <Button colorScheme="teal" variant="solid" size="lg" width={36}>
            退室
          </Button>
        </Grid>
      </Grid>
      <TableContainer
        overflowX="unset"
        overflowY="unset"
        p={14}
        mt={-6}
        outlineOffset={2}
      >
        <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
          <Thead position="sticky" top={0}>
            <Tr bgColor="#E6EBED">
              <Th w="33%">出席者</Th>
              <Th w="33%">部屋</Th>
              <Th w="33%">入室時刻</Th>
            </Tr>
          </Thead>
          <Tbody outline="1px">
            {attendees.map((attendee) => (
              <Tr key={attendee.id}>
                <Td>
                  <Flex alignItems={"center"} gap={3}>
                    <Avatar size={"md"} src={attendee.avaterPath} border="2px">
                      <AvatarBadge boxSize="1.1em" bg="green.500" />
                    </Avatar>
                    {attendee.name}
                  </Flex>
                </Td>
                <Td>{attendee.placeName}</Td>
                <Td>{decodeDate(attendee.enteredAt)}</Td>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </TableContainer>
    </div>
  );
}

export default Home;
