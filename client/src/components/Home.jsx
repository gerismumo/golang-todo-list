import React, { useEffect, useState } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {faTrashCan} from '@fortawesome/free-solid-svg-icons';
import {faPenToSquare } from '@fortawesome/free-regular-svg-icons';
import axios from 'axios';
import { Db_url } from './utils';


const Home = () => {
    const[openEdit, setOpenEdit] = useState(false);
    
    const[todoName, setTodoName] = useState('');
    const[searchTodo, setSearchTodo] = useState('');

    const[todoList, setTodoList] = useState([]);

    const [currentTodo, setCurrentTodo] = useState(null);
    const [editId, setEditId] = useState(null);
    const handleOpenEdit = (id) => {
        setEditId(id);
        setCurrentTodo(todoList.find(todo => todo.id === id));
        console.log(currentTodo);
        if(!openEdit) {
            setOpenEdit(true);
        }else {
            setOpenEdit(false);
        }
    }


    const getTodo = async() => {
        try {
            const response = await axios.get(`${Db_url}/api/getTodo`);
            // console.log(response.data);
            if(response.data.success) {
                setTodoList(response.data.data);
            }
        }catch(error) {
            console.log(error);
        }
    }

    useEffect(() => {
        getTodo();
    },[])

    const handleAddTodo = async(e) => {
        e.preventDefault();

        try {
            const response = await axios.post(`${Db_url}/api/addTodo`, {task:todoName});
            if(response.data.success) {
                getTodo();
            }

        }catch(error) {
            console.log(error.message);
        }

    }

    const handleDeleteTodo = async(id) => {
        try {
            const response = await axios.delete(`${Db_url}/api/deleteTodo/${id}`);
            // console.log(response.data);
            if(response.data.success) {
                getTodo();
            }
        }catch(error) {
            console.log(error.message);
        }
    }


    const handleEditTodo = async(e, id) => {
        e.preventDefault();
        try {
            const response = await axios.put(`${Db_url}/api/editTodo/${id}`, {task: currentTodo.name});
            // console.log(response.data);
            if(response.data.success) {
                getTodo();
                setOpenEdit(false);
            }
        }catch(error) {
            console.log(error.message);
        }
    }

  return (
    <div className="flex flex-row justify-center py-[40px]">
        <div className="flex flex-col justify-center items-center ">
            <div className="">
                <h1 className='text-[35px] font-[700]'>Golang Todo</h1>
            </div>
            <div className="flex flex-row justify-between items-center  gap-[100px] mt-[50px]">
                    <form action="" className='flex flex-row gap-[20px] items-center '>
                        <input type="text" 
                        placeholder="Search Todo" 
                        className=' py-[5px] px-[4px] border-[1px] border-[#068FFF] rounded-[6px] outline-none '
                        />
                        {/* <button type="Search"
                        className='bg-[#4E4FEB] rounded-[6px] text-[#ffffff] py-[5px] px-[20px] '
                        >Search</button> */}
                    </form>
                    <form onSubmit={(e) => handleAddTodo(e)} className='flex flex-row gap-[20px] items-center '>
                        <input type="text" 
                        name='todo'
                        value={todoName}
                        onChange={(e) => setTodoName(e.target.value)}
                        placeholder="Add Todo" 
                        className='py-[5px] px-[4px] border-[1px] border-[#068FFF] rounded-[6px] outline-none '
                        />
                        <button type="submit"
                        className='bg-[#19A7CE] rounded-[6px] text-[#ffffff] py-[5px] px-[20px] '
                        >Add</button>
                    </form>
            </div>
            <div className="mt-[70px]">
                <table className='border-collapse  '>
                    <thead>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Id</th>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Name</th>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Created At</th>
                    </thead>
                    <tbody>
                        {todoList.length === 0 ? (
                            <tr>
                                <td className='border border-[#ddd] px-[30px] py-[10px]'><span>No Data</span></td>
                            </tr>
                        ): (
                            todoList.map((todo) => (
                                <React.Fragment key={todo.id}>
                                    <tr >
                                        <td className='border border-[#ddd] px-[30px] py-[10px]'>{todo.id}</td>
                                        <td className='border border-[#ddd] px-[30px] py-[10px]'>{todo.name}</td>
                                        <td className='border border-[#ddd] px-[30px] py-[10px]'>{todo.createdAt}</td>
                                        <td className='border border-[#ddd] px-[30px] py-[10px]'>
                                            <span onClick={() => handleDeleteTodo(todo.id)}> <FontAwesomeIcon color='red' icon={faTrashCan} /></span>
                                        </td>
                                        <td className='border border-[#ddd] px-[30px] py-[10px]'>
                                            <span onClick={() => handleOpenEdit(todo.id)}><FontAwesomeIcon color='#068FFF' icon={faPenToSquare} /></span>
                                        </td>
                                    </tr>
                                    {openEdit && editId === todo.id && (
                                        <tr>
                                            <td colSpan="5" className='border border-[#ddd] px-[30px] py-[10px]'>
                                                <form onSubmit={(e) => handleEditTodo(e,todo.id)} className='flex flex-row gap-[20px] items-center '>
                                                    <input type="text" 
                                                    name='todo'
                                                    value={currentTodo.name}
                                                    onChange={(e) => setCurrentTodo({...currentTodo,name : e.target.value})}
                                                    placeholder="Edit Todo" 
                                                    className='py-[5px] px-[4px] border-[1px] border-[#068FFF] rounded-[6px] outline-none '
                                                    />
                                                    <button type="submit"
                                                    className='bg-[#19A7CE] rounded-[6px] text-[#ffffff] py-[5px] px-[20px] '
                                                    >Edit</button>
                                                </form>
                                            </td>
                                        </tr>
                                    )}
                                </React.Fragment>
                                
                            ))
                        )}
                    </tbody>
                </table>
            </div>
        </div>
        
    </div>
  )
}

export default Home