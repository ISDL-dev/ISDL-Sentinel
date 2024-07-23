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
import "./Home.css";
import dayjs from "dayjs";
import "dayjs/locale/ja";
import { inRoom } from "../../models/users/user";
import { attendeesListApi } from "../../api";
import { useEffect, useState } from "react";

dayjs.locale("ja");

const decodeDate = (dateString: string) => {
  const date = dayjs(dateString);
  return `${dayjs(date).format("MM月DD日")}（${dayjs(date).format(
    "ddd"
  )}）${dayjs(date).format("HH時MM分")}`;
};

type AuthUser = {
  id: number;
  statusName: string;
};
type Attendee = {
  userId: number;
  userName: string;
  placeName: string;
  statusName: string;
  gradeName: string;
  avaterId: number;
  avaterImgPath: string;
  enteredAt: string;
};
const AUTH_USER: AuthUser = {
  id: 4,
  statusName: inRoom,
};

function Home() {
  const [authUser, setAuthUser] = useState<AuthUser>(AUTH_USER);
  const [attendeeList, setAttendeeList] = useState<Attendee[]>([]);
  const handleStatusChange = async () => {
    try {
      const user = await attendeesListApi.putStatus({
        user_id: authUser.id,
        status: authUser.statusName,
      });
      setAuthUser({ id: user.data.user_id, statusName: user.data.status });
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    (async () => {
      try {
        const response = await attendeesListApi.getAttendeesList();
        console.log(response.data);
        setAttendeeList(
          response.data.map((attendee) => {
            return {
              userId: attendee.user_id,
              userName: attendee.user_name,
              placeName: attendee.place,
              statusName: attendee.status,
              gradeName: attendee.grade,
              avaterId: attendee.avatar_id,
              avaterImgPath: attendee.avatar_img_path,
              enteredAt: attendee.entered_at,
            };
          })
        );
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);
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
          {authUser.statusName === inRoom ? (
            <>
              <Button
                colorScheme="teal"
                variant="solid"
                size="lg"
                width={36}
                isDisabled={true}
              >
                入室済
              </Button>
              <Button
                colorScheme="teal"
                variant="solid"
                size="lg"
                width={36}
                onClick={handleStatusChange}
              >
                退室
              </Button>
            </>
          ) : (
            <>
              <Button
                colorScheme="teal"
                variant="solid"
                size="lg"
                width={36}
                onClick={handleStatusChange}
              >
                入室
              </Button>
              <Button
                colorScheme="teal"
                variant="solid"
                size="lg"
                width={36}
                isDisabled={true}
              >
                退室済
              </Button>
            </>
          )}
        </Grid>
      </Grid>
      <TableContainer
        pb={14}
        pr={14}
        pl={14}
        mt={8}
        outlineOffset={2}
        overflowX="unset"
        overflowY="scroll"
        height="65vh"
      >
        <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
          <Thead top={0}>
            <Tr bgColor="#E6EBED">
              <Th w="33%">出席者</Th>
              <Th w="33%">部屋</Th>
              <Th w="33%">入室時刻</Th>
            </Tr>
          </Thead>
          <Tbody outline="1px">
            {attendeeList.map((attendee) => (
              <Tr key={attendee.userId}>
                <Td>
                  <Flex alignItems={"center"} gap={3}>
                    <Avatar
                      size={"md"}
                      src={`./avatar/${attendee.avaterImgPath}`}
                      border="2px"
                    ></Avatar>
                    {attendee.userName}
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
