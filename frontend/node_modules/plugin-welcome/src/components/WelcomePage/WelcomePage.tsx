import React, { FC, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import {useLocation} from "react-router-dom";

 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบผู้เสียชีวิต' };
  const username = { givenuser: 'นายแพทย์กำลัง นอนพอดี' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName}`}
       subtitle=""
     ></Header>
     <Content>
       <ContentHeader title="ตารางข้อมูลผู้เสียชีวิต">
         <div>
         <Button variant="contained" color="primary">
             ออกจากระบบ
         </Button>
         </div>
       </ContentHeader>
       <ContentHeader title="">
         <div>
          <Link component={RouterLink} to="/deceasedreceive">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              เพิ่มข้อมูลผู้เสียชีวิต
            </Button>
          </Link>
         </div>
       </ContentHeader>
       <div>
       </div>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;