import {
  Avatar,
  AvatarBadge,
  Button,
  Flex,
  Grid,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Table,
  TableContainer,
  Tabs,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from "@chakra-ui/react";
import { useState } from "react";
import { RankingList } from "../../features/RankingList";

export const Ranking = () => {
  const colors = useColorModeValue(
    ["red.50", "blue.50"],
    ["red.900", "blue.900"]
  );
  const [tabIndex, setTabIndex] = useState(0);
  const bg = colors[tabIndex];
  return (
    <Tabs
      isFitted
      onChange={(index) => setTabIndex(index)}
      bg={bg}
      border={"2px"}
      borderColor="gray.200"
      mt={{ base: 20, md: 0 }}
      width="100%"
    >
      <TabList mb="1em">
        <Tab>滞在時間</Tab>
        <Tab>出席日数</Tab>
      </TabList>
      <TabPanels>
        <TabPanel>
          <RankingList placeholder="滞在時間"></RankingList>
        </TabPanel>
        <TabPanel>
          <RankingList placeholder="出席日数"></RankingList>
        </TabPanel>
      </TabPanels>
    </Tabs>
  );
};
