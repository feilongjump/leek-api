<?php

namespace App\Models;

use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * @property string $markdown
 * @property string $body
 */
class Content extends Model
{
    use SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var string[]
     */
    protected $fillable = [
        'contentable_type', 'contentable_id', 'body', 'markdown',
    ];

    public function contentable()
    {
        return $this->morphTo();
    }
}
