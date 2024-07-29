import { useEffect, useState } from "react";
import {
  Grid,
  Table,
  TableContainer,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Flex,
  Avatar,
  chakra,
  Box,
  Text,
} from "@chakra-ui/react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./access_history.css";
import { accessHistoryApi } from "../../api";
import { GetAccessHistory200ResponseInner } from "../../schema";
import { useNavigate } from "react-router-dom";

export default function AccessHistory() {
  const [accessHistory, setAccessHistoryData] = useState<
    GetAccessHistory200ResponseInner[]
  >([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const navigate = useNavigate();

  const handleDateChange = (date: Date | null) => {
    if (date) {
      setSelectedDate(date);
    }
  };

  const formatDate = (date: Date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    return `${year}-${month}`;
  };

  const formatTime = (dateTime: string) => {
    const date = new Date(dateTime);
    return date.toISOString().substring(11, 19);
  };

  useEffect(() => {
    const fetchUserData = async (date: Date) => {
      setLoading(true);
      try {
        const formattedDate = formatDate(date);
        const response = await accessHistoryApi.getAccessHistory(formattedDate);
        setAccessHistoryData(response.data ?? []);
      } catch (err) {
        setError("データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchUserData(selectedDate);
  }, [selectedDate]); 

  useEffect(() => {
    const now = new Date();
    const currentMonth = new Date(now.getFullYear(), now.getMonth(), 1);
    setSelectedDate(currentMonth);
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;

  return (
    <Box p={6} mt={{ base: 20, md: 0 }}>
      <Grid
        templateColumns={{ base: "1fr", md: "1fr auto" }}
        alignItems="center"
      >
        <Text
          fontSize={{ base: "2xl", md: "4xl" }}
          textAlign="left"
          className="block font-bold text-gray-900 dark:text-white"
        >
          入退室履歴
        </Text>
        <Flex
          justifyContent={{ base: "center", md: "flex-end" }}
          alignItems="center"
          mt={{ base: 4, md: 0 }}
          flexDirection={{ base: "column", md: "row" }}
          textAlign={{ base: "center", md: "left" }}
        >
          <chakra.label htmlFor="month-picker" mr={3} fontWeight="bold">
            月を選択:
          </chakra.label>
          <DatePicker
            selected={selectedDate}
            onChange={handleDateChange}
            dateFormat="yyyy/MM"
            showMonthYearPicker
            id="month-picker"
            className="custom-datepicker"
          />
        </Flex>
      </Grid>
      <TableContainer
        pb={14}
        pr={{ base: 2, md: 14 }}
        pl={{ base: 2, md: 14 }}
        mt={8}
        outlineOffset={2}
        overflowX="unset"
        overflowY="scroll"
        height="65vh"
      >
        <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
          <Thead top={0}>
            <Tr bgColor="#E6EBED">
              <Th w="20%">日付</Th>
              <Th w="20%">入室者</Th>
              <Th w="20%">入室時間</Th>
              <Th w="20%">退室者</Th>
              <Th w="20%">退室時間</Th>
            </Tr>
          </Thead>
          <Tbody outline="1px">
            {accessHistory.length === 0 ? (
              <Tr>
                <Td colSpan={5} textAlign="center">
                  データがありません
                </Td>
              </Tr>
            ) : (
              accessHistory.map((access, index) => (
                <Tr key={index}>
                  <Td>{access.date}</Td>
                  <Td>
                    <Flex alignItems={"center"} gap={3}>
                      {access.entering.avatar_img_path && (
                        <Avatar
                          size={"md"}
                          src={`./avatar/${access.entering.avatar_img_path}`}
                          border="2px"
                          onClick={() =>
                            navigate("/profile", {
                              state: { userId: access.entering.user_id },
                            })
                          }
                        />
                      )}
                      {access.entering.user_name}
                    </Flex>
                  </Td>
                  <Td>{formatTime(access.entering.entered_at)}</Td>
                  <Td>
                    {access.leaving.left_at &&
                    access.leaving.avatar_img_path ? (
                      <Flex alignItems={"center"} gap={3}>
                        <Avatar
                          size={"md"}
                          src={`./avatar/${access.leaving.avatar_img_path}`}
                          border="2px"
                          onClick={() =>
                            navigate("/profile", {
                              state: { userId: access.leaving.user_id },
                            })
                          }
                        />
                        {access.leaving.user_name}
                      </Flex>
                    ) : null}
                  </Td>
                  <Td>
                    {access.leaving.left_at
                      ? formatTime(access.leaving.left_at)
                      : null}
                  </Td>
                </Tr>
              ))
            )}
          </Tbody>
        </Table>
      </TableContainer>
    </Box>
  );
}
