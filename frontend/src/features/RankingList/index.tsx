import {
  Card,
  CardBody,
  Grid,
  GridItem,
  Heading,
  Image,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
} from "@chakra-ui/react";
import { Top3Icon } from "../RankIcon/Top3Icon";
import { Top10Icon } from "../RankIcon/Top10Icon";
import { useEffect, useState } from "react";
import { GetRanking200ResponseInner } from "../../schema";
import { useNavigate } from "react-router-dom";

interface RankingListProps {
  getRankList: GetRanking200ResponseInner[];
  placeholder: string;
}

export const RankingList: React.FC<RankingListProps> = ({
  getRankList,
  placeholder,
}) => {
  type RankedItem = GetRanking200ResponseInner & { rank: number };
  const [rankingList, setRankingList] = useState<RankedItem[]>([]);
  const navigate = useNavigate();
  const getDateFromStringFormat = (dateFormat: string) => {
    const [hours, minutes, seconds] = dateFormat.split(":").map(Number);
    const totalSeconds = hours * 3600 + minutes * 60 + seconds;
    const baseDate = new Date(0);
    return new Date(baseDate.getTime() + totalSeconds * 1000);
  };
  const descTimeSort = (a: Date, b: Date) => {
    return a < b ? 1 : -1;
  };
  const orderByPlaceholder = (
    responseData: GetRanking200ResponseInner[],
    placeholder: string
  ) => {
    return placeholder === "stay_time"
      ? responseData.sort((a, b) =>
          descTimeSort(
            getDateFromStringFormat(a.stay_time),
            getDateFromStringFormat(b.stay_time)
          )
        )
      : responseData.sort((a, b) => b.attendance_days - a.attendance_days);
  };
  const assignRanks = (
    list: GetRanking200ResponseInner[],
    placeholder: string
  ): RankedItem[] => {
    const sortedList = orderByPlaceholder([...list], placeholder);
    let currentRank = 1;
    let previousValue: number | string =
      placeholder === "stay_time"
        ? sortedList[0].stay_time
        : sortedList[0].attendance_days;
    return sortedList.map((item, index) => {
      const currentValue =
        placeholder === "stay_time" ? item.stay_time : item.attendance_days;
      if (index > 0 && currentValue !== previousValue) {
        currentRank = index + 1;
      }
      previousValue = currentValue;
      return { ...item, rank: currentRank };
    });
  };
  const formatResultByPlaceholder = (item: GetRanking200ResponseInner) => {
    return placeholder === "stay_time"
      ? item.stay_time
      : `${item.attendance_days}日`;
  };
  const LIMIT_DISPLAY_RANK = 10;
  const awardUserIndex = 3;
  const endIndex = LIMIT_DISPLAY_RANK ?? rankingList.length;
  useEffect(() => {
    const rankedList = assignRanks(getRankList, placeholder);
    setRankingList(rankedList as RankedItem[]);
  }, [getRankList, placeholder]);
  return (
    <div>
      <Grid
        templateColumns="repeat(3, 1fr)"
        templateRows="1fr 2fr 1fr 1fr 12fr"
        justifyItems={"center"}
        placeItems={"center"}
        alignItems={"center"}
        height={"20%"}
        w={"-moz-max-content"}
        column={3}
        row={4}
      >
        {rankingList.length !== 0 &&
          rankingList.slice(0, awardUserIndex).map((item, index) => (
            <GridItem
              rowSpan={4}
              colSpan={1}
              textAlign="center"
              key={item.user_id}
              display="flex"
              flexDirection="column"
              justifyContent="center"
              alignItems="center"
            >
              <Top3Icon rank={item.rank - 1}></Top3Icon>
              <Image
                src={item.avatar_img_path}
                alt={`${item.avatar_id}`}
                boxSize={{
                  base: "48px",
                  md: "64px",
                }}
                cursor="pointer"
                onClick={() =>
                  navigate("/profile", {
                    state: { userId: item.user_id },
                  })
                }
              />
              <Text>{item.user_name}</Text>
              <Heading fontSize={{ base: "xl", md: "2xl" }} color="#fa8072">
                {formatResultByPlaceholder(item)}
              </Heading>
            </GridItem>
          ))}

        <GridItem rowSpan={1} colSpan={3}>
          <TableContainer
            pb={14}
            pr={14}
            pl={14}
            pt={3}
            overflowX="auto"
            overflowY="scroll"
            height={{
              base: "47vh",
              md: "48vh",
            }}
            width={{
              base: "120vw",
              md: "60vw",
            }}
          >
            <Table size="sm">
              <Tbody>
                {rankingList.length > 0 &&
                  rankingList.slice(awardUserIndex, endIndex).map((item) => (
                    <Tr key={item.user_id}>
                      <Td textAlign="center" w={15}>
                        <Card>
                          <CardBody>
                            <Grid
                              templateColumns="auto auto 1fr auto"
                              alignItems={"center"}
                              height={"20%"}
                              w={"-moz-max-content"}
                              gap={4}
                            >
                              <GridItem
                                justifySelf="start"
                                marginRight={{ base: 12, md: 20 }}
                                paddingTop={14}
                              >
                                <Top10Icon rank={`${item.rank}`} />
                              </GridItem>
                              <GridItem
                                justifySelf="start"
                                marginRight={{ base: 0, md: 8 }}
                              >
                                <Image
                                  src={item.avatar_img_path}
                                  alt={`${item.avatar_id}`}
                                  boxSize={{
                                    base: "36px",
                                    md: "50px",
                                  }}
                                  cursor="pointer"
                                  onClick={() =>
                                    navigate("/profile", {
                                      state: {
                                        userId: item.user_id,
                                      },
                                    })
                                  }
                                />
                              </GridItem>
                              <GridItem justifySelf="start">
                                <Text fontSize="large">{item.user_name}</Text>
                              </GridItem>
                              <GridItem justifySelf="end">
                                <Heading
                                  fontSize={{ base: "xl", md: "2xl" }}
                                  color="#fa8072"
                                >
                                  {formatResultByPlaceholder(item)}
                                </Heading>
                              </GridItem>
                            </Grid>
                          </CardBody>
                        </Card>
                      </Td>
                    </Tr>
                  ))}
              </Tbody>
            </Table>
          </TableContainer>
        </GridItem>
      </Grid>
    </div>
  );
};
