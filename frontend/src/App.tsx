import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import {
  ChakraProvider,
  Box,
  useColorModeValue,
  Drawer,
  DrawerContent,
} from "@chakra-ui/react";
import Home from "./routes/Home";
import AccessHistory from "./routes/AccessHistory";
import Profile from "./routes/Profile";
import SignInWebauthn from "./routes/SignInWebauthn";
import Footer from "./features/Footer";
import SidebarContent from "./features/SidebarContent";
import MobileNav from "./features/MobileNav";
import { useDisclosure } from "@chakra-ui/react";
import { Ranking } from "./routes/Ranking";

const App = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <ChakraProvider>
      <BrowserRouter>
        <Box minH="100vh" bg={useColorModeValue("gray.100", "gray.900")}>
          <MobileNav onOpen={onOpen} />
          <SidebarContent
            onClose={onClose}
            display={{ base: "none", md: "block" }}
            pt={20}
          />
          <Drawer
            isOpen={isOpen}
            placement="left"
            onClose={onClose}
            returnFocusOnClose={false}
            onOverlayClick={onClose}
            size="full"
          >
            <DrawerContent>
              <SidebarContent onClose={onClose} />
            </DrawerContent>
          </Drawer>
          <Box
            pl={{ base: 2, md: 64 }}
            pr={{ base: 2, md: 6 }}
            pt={{ base: 2, md: 24 }}
          >
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/access-history" element={<AccessHistory />} />
              <Route path="/profile" element={<Profile />} />
              <Route path="/ranking" element={<Ranking />} />
              <Route path="/sign-in-webauthn" element={<SignInWebauthn />} />
            </Routes>
          </Box>
        </Box>
        <Footer />
      </BrowserRouter>
    </ChakraProvider>
  );
};

export default App;
