import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntCoolroom } from '../../api/models/EntCoolroom';
import { EntCoolroomType } from '../../api/models/EntCoolroomType';
import { EntRelative } from '../../api/models/EntRelative';
import { EntPatient } from '../../api/models/EntPatient';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);


export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบผู้เสียชีวิต' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [coolrooms, setCoolrooms] = useState<EntCoolroom[]>([]);
  const [coolroomtypes, setCoolroomTypes] = useState<EntCoolroomType[]>([]);
  const [relatives, setRelatives] = useState<EntRelative[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [loading, setLoading] = useState(true);
  const [deathtime, setDeathtime] = useState(String);

  const [coolroomid, setcoolroom] = useState(Number);
  const [coolroomtypeid, setCoolroomType] = useState(Number);
  const [relativeid, setRelative] = useState(Number);
  const [patientid, setPatient] = useState(Number);
  const [userid, setUser] = useState(Number);

  useEffect(() => {
    const getCoolroomTypes = async () => {
      const res = await api.listCoolroomtype({ limit: 10, offset: 0 });
      setLoading(false);
      setCoolroomTypes(res);
    };
    getCoolroomTypes();

    const getRelatives = async () => {
      const res = await api.listRelative({ limit: 10, offset: 0 });
      setLoading(false);
      setRelatives(res);
    };
    getRelatives();

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
    };
    getPatients();

    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();

  }, [loading]);


  const DeathTimehandleChange = (event: any) => {
    setDeathtime(event.target.value as string);
  };

  const CoolroomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setcoolroom(event.target.value as number);
  };

  const CoolroomTypehandleChange = async (event: React.ChangeEvent<{ value: unknown }>) => {
    setCoolroomType(event.target.value as number);
    const res = await api.listCoolroomByCoolroomtype({ typeid: event.target.value as number, limit: 10, offset: 0 });
    setCoolrooms(res);
  };

  const RelativehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRelative(event.target.value as number);
  };

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const UserthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const CreateDeceasedReceive = async () => {
    const deceasedreceive = {
      patientID: patientid,
      coolroomID: coolroomid,
      relativeID: relativeid,
      userID: userid,
      deathtime: deathtime + ":00+07:00"

    };
    console.log(deceasedreceive);
    const res: any = await api.createDeceasedreceive({ deceasedreceive: deceasedreceive });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลผู้เสียชีวิต">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    พบปัญหาระหว่างบันทึกข้อมูล
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="user-label">แพทย์</InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userid}
                  onChange={UserthandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.userName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="patient-label">ชื่อ-สกุล ผู้เสียชีวิต</InputLabel>
              <Select
                labelId="patient-label"
                id="patient"
                value={patientid}
                onChange={PatienthandleChange}
                style={{ width: 400 }}
              >
                {patients.filter((filter: any) => filter.edges.deceasedreceives == null).map((item: EntPatient) => (
                  <MenuItem value={item.id}>{item.patientName}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="relative-label">ชื่อ-สกุล ญาติผู้เสียชีวิต</InputLabel>
                <Select
                  labelId="relative-label"
                  id="relative"
                  value={relativeid}
                  onChange={RelativehandleChange}
                  style={{ width: 400 }}
                >
                  {relatives.map((item: any) => (
                    <MenuItem value={item.id}>{item.relativeName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="status-death">ประเภทห้องเย็น</InputLabel>
              <Select
                labelId="status-death"
                id="statusdeath"
                value={coolroomtypeid}
                onChange={CoolroomTypehandleChange}
                style={{ width: 200 }}
              >
                {coolroomtypes.map((item: EntCoolroomType) => (
                  <MenuItem key={item.id} value={item.id}>{item.coolroomtypeName}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="coolroom-label">หมายเลขห้องเย็น</InputLabel>
                <Select
                  labelId="coolroom-label"
                  id="coolroom"
                  value={coolroomid}
                  onChange={CoolroomhandleChange}
                  style={{ width: 200 }}
                >
                  {coolrooms.map(item => (
                    <MenuItem value={item.id}>{item.coolroomName}</MenuItem>
                  ))}
                </Select>
              </FormControl>

              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="deathtime"
                  label="เดือน/วัน/ปี เวลาเสียชีวิต"
                  type="datetime-local"
                  value={deathtime}
                  onChange={DeathTimehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateDeceasedReceive();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
