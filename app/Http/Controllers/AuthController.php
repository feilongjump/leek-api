<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Requests\AuthRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Contracts\Auth\Authenticatable;

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

    /**
     * Get the token array structure.
     *
     * @param User|Authenticatable $user
     * @param array $abilities
     * @return \Illuminate\Http\JsonResponse
     */
    protected function respondWithToken(User|Authenticatable $user, array $abilities = ['*']): \Illuminate\Http\JsonResponse
    {
        return response()->json([
            'access_token' => $user->createToken(config('app.name'), $abilities)->plainTextToken,
            'token_type' => 'Bearer',
            'expires_in' => now()->addSeconds(config('sanctum.expiration') ?? 604800)->timestamp
        ]);
    }
}
