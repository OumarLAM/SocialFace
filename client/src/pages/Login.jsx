import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useForm } from "react-hook-form";
import { TbSocial } from "react-icons/tb";
import { useDispatch } from "react-redux";
import { CustomButton, Loading, TextInput } from "../components";
import { BackgroundImage } from "../assets";

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    mode: "onChange",
  });

  const [errorMessage, setErrorMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const dispatch = useDispatch();

  return (
    <div className="bg-bgColor w-full h-[100vh] flex items-center justify-center p-6">
      <div className="w-full md:w-2/3 h-fit lg:h-full 2xl:h-5/6 py-8 lg:py-0 flex bg-primary rounded-xl overflow-hidden shadow-xl"></div>
      <div className="w-full lg:w=1/2 h-full p-10 2xl:px-20 flex flex-col justify-center">
        <div className="w-full flex gap-2 items-center mb-6">
          <div className="p-2 bg-[#065ad8] rounded text-white">
            <TbSocial />
          </div>
          <span className="text-2xl text-[#065ad8]">SocialFace</span>
        </div>

        <p className="text-ascent-1 text-base font-semibold">
          Login to your account
        </p>
        <span className="text-sm mt-2 text-ascent-2">Welcome back</span>

        <form className="py-8 flex flex-col gap-5">
          <TextInput
            name="email"
            placeholder="example@gmail.com"
            label="Email Address"
            type="email"
            register={register("email", {
              required: "Email Address is required",
            })}
            styles="w-full rounded-full"
            labelStyle="ml-2"
            error={errors.email ? errors.email.message : ""}
          />

          <TextInput
            name="password"
            placeholder="Password"
            label="Password"
            type="password"
            styles="w-full rounded-full"
            labelStyle="ml-2"
            register={register("password", {
              required: "Password is required",
            })}
            error={errors.password ? errors.password?.message : ""}
          />

          {errorMessage?.message && (
            <span
              className={`text-sm ${
                errorMessage?.status === "failed"
                  ? "text-[#f64949fe]"
                  : "text-[#2ba150fe]"
              } mt-0.5`}
            >
              {errorMessage?.message}
            </span>
          )}

          {isSubmitting ? (
            <Loading />
          ) : (
            <CustomButton
              type="submit"
              containerStyle={`inline-flex justify-center rounded-md bg-blue px-8 py-3 text-sm font-medium text-white outline-none`}
              title="Login"
            />
          )}
        </form>

        <p className="text-ascent-2 text-sm text-center">
          Don't have an account ?
          <Link
            to="/register"
            className="text-[#065ad8] font-semibold ml-2 cursor-pointer"
          >
            Create a new account
          </Link>
        </p>
      </div>
      <div className="hidden w-1/2 h-full lg:flex flex-col items-center justify-center bg-blue">
        <div className="relative w-full flex items-center justify-center">
          <img
            src={BackgroundImage}
            alt="Background Image"
            className="w-48 2xl:w-64 h-48 2xl:h-64 rounded-full object-cover"
          />
        </div>
      </div>
    </div>
  );
};

export default Login;
