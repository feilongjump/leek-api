<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Auth\Events\PasswordReset;
use Illuminate\Http\Request;
use App\Http\Requests\AuthRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Contracts\Auth\Authenticatable;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Password;
use Illuminate\Support\Str;
use Symfony\Component\HttpFoundation\Response;

class AuthController extends Controller
{
    public function login(AuthRequest $request)
    {
        $username = $request->username;

        filter_var($username, FILTER_VALIDATE_EMAIL) ?
            $credentials['email'] = $username :
            $credentials['name'] = $username;

        $credentials['password'] = $request->password;

        if (!Auth::attempt($credentials)) {
            throw new AuthenticationException('用户名或密码错误');
        }

        $user = Auth::user();
        $abilities = $user->isActivated ? ['*'] : [];

        return $this->respondWithToken($user, $abilities)->setStatusCode(201);
    }

    public function reset(Request $request)
    {
        return $request->get('email') ?
            $this->resetPasswordByToken($request) :
            $this->resetPassword($request);
    }

    public function resetPasswordByToken(Request $request)
    {
        $request->validate([
            'token' => 'required',
            'email' => 'required|email',
            'password' => 'required|confirmed|min:6',
        ]);

        $status = $this->broker()->reset(
            $request->only('email', 'password', 'password_confirmation', 'token'),
            function ($user, $password) {

                $user->forceFill([
                    'password' => Hash::make($password)
                ])->setRememberToken(Str::random(60));

                $user->save();

                event(new PasswordReset($user));
            }
        );

        return $status === Password::PASSWORD_RESET ?
            response()->json($this->resetPasswordMessage($status)) :
            response()->json($this->resetPasswordMessage($status))->setStatusCode(Response::HTTP_UNAUTHORIZED);
    }

    public function resetPassword(Request $request)
    {
        $request->validate([
            'old_password' => 'required|current_password:sanctum',
            'password' => 'required|different:old_password|confirmed|min:6',
        ], [
            'old_password.current_password' => '旧密码输入错误！'
        ], [
            'old_password' => '旧密码',
        ]);

        auth()->guard('sanctum')->user()->update([
            'password' => bcrypt($request->get('password')),
        ]);

        return response()->json($this->resetPasswordMessage($this->broker()::PASSWORD_RESET));
    }

    public function forgetPassword(AuthRequest $request)
    {
        $this->broker()->sendResetLink(
            $request->only('email')
        );

        return response()->json($this->resetPasswordMessage($this->broker()::RESET_LINK_SENT));
    }

    /**
     * Get the token array structure.
     *
     * @param User|Authenticatable $user
     * @param array $abilities
     * @return \Illuminate\Http\JsonResponse
     */
    protected function respondWithToken(User|Authenticatable $user, array $abilities = ['*']): \Illuminate\Http\JsonResponse
    {
        $expiredAt = config('sanctum.expiration');

        return response()->json([
            'access_token' => $user->createToken(config('app.name'), $abilities)->plainTextToken,
            'token_type' => 'Bearer',
            'expires_in' => now()->addMinutes($expiredAt)->timestamp
        ]);
    }

    protected function broker()
    {
        return Password::broker();
    }

    protected function resetPasswordMessage(string $status): array
    {
        $message = __($status);

        return compact('message');
    }
}
